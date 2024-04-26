package coinbasemanager

import (
	"math"

	"github.com/pkg/errors"
	"github.com/spectre-project/spectred/domain/consensus/model"
	"github.com/spectre-project/spectred/domain/consensus/model/externalapi"
	"github.com/spectre-project/spectred/domain/consensus/utils/constants"
	"github.com/spectre-project/spectred/domain/consensus/utils/hashset"
	"github.com/spectre-project/spectred/domain/consensus/utils/subnetworks"
	"github.com/spectre-project/spectred/domain/consensus/utils/transactionhelper"
	"github.com/spectre-project/spectred/infrastructure/db/database"
)

type coinbaseManager struct {
	subsidyGenesisReward                    uint64
	preDeflationaryPhaseBaseSubsidy         uint64
	coinbasePayloadScriptPublicKeyMaxLength uint8
	genesisHash                             *externalapi.DomainHash
	deflationaryPhaseDaaScore               uint64
	deflationaryPhaseBaseSubsidy            uint64

	databaseContext     model.DBReader
	dagTraversalManager model.DAGTraversalManager
	ghostdagDataStore   model.GHOSTDAGDataStore
	acceptanceDataStore model.AcceptanceDataStore
	daaBlocksStore      model.DAABlocksStore
	blockStore          model.BlockStore
	pruningStore        model.PruningStore
	blockHeaderStore    model.BlockHeaderStore
}

func (c *coinbaseManager) ExpectedCoinbaseTransaction(stagingArea *model.StagingArea, blockHash *externalapi.DomainHash,
	coinbaseData *externalapi.DomainCoinbaseData) (expectedTransaction *externalapi.DomainTransaction, hasRedReward bool, err error) {

	ghostdagData, err := c.ghostdagDataStore.Get(c.databaseContext, stagingArea, blockHash, true)
	if !database.IsNotFoundError(err) && err != nil {
		return nil, false, err
	}

	// If there's ghostdag data with trusted data we prefer it because we need the original merge set non-pruned merge set.
	if database.IsNotFoundError(err) {
		ghostdagData, err = c.ghostdagDataStore.Get(c.databaseContext, stagingArea, blockHash, false)
		if err != nil {
			return nil, false, err
		}
	}

	acceptanceData, err := c.acceptanceDataStore.Get(c.databaseContext, stagingArea, blockHash)
	if err != nil {
		return nil, false, err
	}

	daaAddedBlocksSet, err := c.daaAddedBlocksSet(stagingArea, blockHash)
	if err != nil {
		return nil, false, err
	}

	txOuts := make([]*externalapi.DomainTransactionOutput, 0, len(ghostdagData.MergeSetBlues()))
	acceptanceDataMap := acceptanceDataFromArrayToMap(acceptanceData)
	for _, blue := range ghostdagData.MergeSetBlues() {
		txOut, hasReward, err := c.coinbaseOutputForBlueBlock(stagingArea, blue, acceptanceDataMap[*blue], daaAddedBlocksSet)
		if err != nil {
			return nil, false, err
		}

		if hasReward {
			txOuts = append(txOuts, txOut)
		}
	}

	txOut, hasRedReward, err := c.coinbaseOutputForRewardFromRedBlocks(
		stagingArea, ghostdagData, acceptanceData, daaAddedBlocksSet, coinbaseData)
	if err != nil {
		return nil, false, err
	}

	if hasRedReward {
		txOuts = append(txOuts, txOut)
	}

	subsidy, err := c.CalcBlockSubsidy(stagingArea, blockHash)
	if err != nil {
		return nil, false, err
	}

	payload, err := c.serializeCoinbasePayload(ghostdagData.BlueScore(), coinbaseData, subsidy)
	if err != nil {
		return nil, false, err
	}

	return &externalapi.DomainTransaction{
		Version:      constants.MaxTransactionVersion,
		Inputs:       []*externalapi.DomainTransactionInput{},
		Outputs:      txOuts,
		LockTime:     0,
		SubnetworkID: subnetworks.SubnetworkIDCoinbase,
		Gas:          0,
		Payload:      payload,
	}, hasRedReward, nil
}

func (c *coinbaseManager) daaAddedBlocksSet(stagingArea *model.StagingArea, blockHash *externalapi.DomainHash) (
	hashset.HashSet, error) {

	daaAddedBlocks, err := c.daaBlocksStore.DAAAddedBlocks(c.databaseContext, stagingArea, blockHash)
	if err != nil {
		return nil, err
	}

	return hashset.NewFromSlice(daaAddedBlocks...), nil
}

// coinbaseOutputForBlueBlock calculates the output that should go into the coinbase transaction of blueBlock
// If blueBlock gets no fee - returns nil for txOut
func (c *coinbaseManager) coinbaseOutputForBlueBlock(stagingArea *model.StagingArea,
	blueBlock *externalapi.DomainHash, blockAcceptanceData *externalapi.BlockAcceptanceData,
	mergingBlockDAAAddedBlocksSet hashset.HashSet) (*externalapi.DomainTransactionOutput, bool, error) {

	blockReward, err := c.calcMergedBlockReward(stagingArea, blueBlock, blockAcceptanceData, mergingBlockDAAAddedBlocksSet)
	if err != nil {
		return nil, false, err
	}

	if blockReward == 0 {
		return nil, false, nil
	}

	// the ScriptPublicKey for the coinbase is parsed from the coinbase payload
	_, coinbaseData, _, err := c.ExtractCoinbaseDataBlueScoreAndSubsidy(blockAcceptanceData.TransactionAcceptanceData[0].Transaction)
	if err != nil {
		return nil, false, err
	}

	txOut := &externalapi.DomainTransactionOutput{
		Value:           blockReward,
		ScriptPublicKey: coinbaseData.ScriptPublicKey,
	}

	return txOut, true, nil
}

func (c *coinbaseManager) coinbaseOutputForRewardFromRedBlocks(stagingArea *model.StagingArea,
	ghostdagData *externalapi.BlockGHOSTDAGData, acceptanceData externalapi.AcceptanceData, daaAddedBlocksSet hashset.HashSet,
	coinbaseData *externalapi.DomainCoinbaseData) (*externalapi.DomainTransactionOutput, bool, error) {

	acceptanceDataMap := acceptanceDataFromArrayToMap(acceptanceData)
	totalReward := uint64(0)
	for _, red := range ghostdagData.MergeSetReds() {
		reward, err := c.calcMergedBlockReward(stagingArea, red, acceptanceDataMap[*red], daaAddedBlocksSet)
		if err != nil {
			return nil, false, err
		}

		totalReward += reward
	}

	if totalReward == 0 {
		return nil, false, nil
	}

	return &externalapi.DomainTransactionOutput{
		Value:           totalReward,
		ScriptPublicKey: coinbaseData.ScriptPublicKey,
	}, true, nil
}

func acceptanceDataFromArrayToMap(acceptanceData externalapi.AcceptanceData) map[externalapi.DomainHash]*externalapi.BlockAcceptanceData {
	acceptanceDataMap := make(map[externalapi.DomainHash]*externalapi.BlockAcceptanceData, len(acceptanceData))
	for _, blockAcceptanceData := range acceptanceData {
		acceptanceDataMap[*blockAcceptanceData.BlockHash] = blockAcceptanceData
	}
	return acceptanceDataMap
}

// CalcBlockSubsidy returns the subsidy amount a block at the provided blue score
// should have. This is mainly used for determining how much the coinbase for
// newly generated blocks awards as well as validating the coinbase for blocks
// has the expected value.
func (c *coinbaseManager) CalcBlockSubsidy(stagingArea *model.StagingArea, blockHash *externalapi.DomainHash) (uint64, error) {
	if blockHash.Equal(c.genesisHash) {
		return c.subsidyGenesisReward, nil
	}
	blockDaaScore, err := c.daaBlocksStore.DAAScore(c.databaseContext, stagingArea, blockHash)
	if err != nil {
		return 0, err
	}
	if blockDaaScore < c.deflationaryPhaseDaaScore {
		return c.preDeflationaryPhaseBaseSubsidy, nil
	}

	blockSubsidy := c.calcDeflationaryPeriodBlockSubsidy(blockDaaScore)
	return blockSubsidy, nil
}

func (c *coinbaseManager) calcDeflationaryPeriodBlockSubsidy(blockDaaScore uint64) uint64 {
	// We define a year as 365.25 days and a month as 365.25 / 12 = 30.4375
	// secondsPerMonth = 30.4375 * 24 * 60 * 60
	const secondsPerMonth = 2629800
	// Note that this calculation implicitly assumes that block per second = 1 (by assuming daa score diff is in second units).
	monthsSinceDeflationaryPhaseStarted := (blockDaaScore - c.deflationaryPhaseDaaScore) / secondsPerMonth
	// Return the pre-calculated value from subsidy-per-month table
	return c.getDeflationaryPeriodBlockSubsidyFromTable(monthsSinceDeflationaryPhaseStarted)
}

/*
This table was pre-calculated by calling `calcDeflationaryPeriodBlockSubsidyFloatCalc` for all months until reaching 0 subsidy.
To regenerate this table, run `TestBuildSubsidyTable` in coinbasemanager_test.go (note the `deflationaryPhaseBaseSubsidy` therein)
*/
var subsidyByDeflationaryMonthTable = []uint64{
	1200000000, 1175000000, 1150000000, 1125000000, 1100000000, 1075000000, 1050000000, 1025000000, 1000000000, 975000000, 950000000, 925000000, 900000000, 875000000, 850000000, 825000000, 800000000, 775000000, 750000000, 725000000, 700000000, 675000000, 650000000, 625000000, 600000000,
	587500000, 575000000, 562500000, 550000000, 537500000, 525000000, 512500000, 500000000, 487500000, 475000000, 462500000, 450000000, 437500000, 425000000, 412500000, 400000000, 387500000, 375000000, 362500000, 350000000, 337500000, 325000000, 312500000, 300000000, 293750000,
	287500000, 281250000, 275000000, 268750000, 262500000, 256250000, 250000000, 243750000, 237500000, 231250000, 225000000, 218750000, 212500000, 206250000, 200000000, 193750000, 187500000, 181250000, 175000000, 168750000, 162500000, 156250000, 150000000, 146875000, 143750000,
	140625000, 137500000, 134375000, 131250000, 128125000, 125000000, 121875000, 118750000, 115625000, 112500000, 109375000, 106250000, 103125000, 100000000, 96875000, 93750000, 90625000, 87500000, 84375000, 81250000, 78125000, 75000000, 73437500, 71875000, 70312500,
	68750000, 67187500, 65625000, 64062500, 62500000, 60937500, 59375000, 57812500, 56250000, 54687500, 53125000, 51562500, 50000000, 48437500, 46875000, 45312500, 43750000, 42187500, 40625000, 39062500, 37500000, 36718750, 35937500, 35156250, 34375000,
	33593750, 32812500, 32031250, 31250000, 30468750, 29687500, 28906250, 28125000, 27343750, 26562500, 25781250, 25000000, 24218750, 23437500, 22656250, 21875000, 21093750, 20312500, 19531250, 18750000, 18359375, 17968750, 17578125, 17187500, 16796875,
	16406250, 16015625, 15625000, 15234375, 14843750, 14453125, 14062500, 13671875, 13281250, 12890625, 12500000, 12109375, 11718750, 11328125, 10937500, 10546875, 10156250, 9765625, 9375000, 9179687, 8984375, 8789062, 8593750, 8398437, 8203125,
	8007812, 7812500, 7617187, 7421875, 7226562, 7031250, 6835937, 6640625, 6445312, 6250000, 6054687, 5859375, 5664062, 5468750, 5273437, 5078125, 4882812, 4687500, 4589843, 4492187, 4394531, 4296875, 4199218, 4101562, 4003906,
	3906250, 3808593, 3710937, 3613281, 3515625, 3417968, 3320312, 3222656, 3125000, 3027343, 2929687, 2832031, 2734375, 2636718, 2539062, 2441406, 2343750, 2294921, 2246093, 2197265, 2148437, 2099609, 2050781, 2001953, 1953125,
	1904296, 1855468, 1806640, 1757812, 1708984, 1660156, 1611328, 1562500, 1513671, 1464843, 1416015, 1367187, 1318359, 1269531, 1220703, 1171875, 1147460, 1123046, 1098632, 1074218, 1049804, 1025390, 1000976, 976562, 952148,
	927734, 903320, 878906, 854492, 830078, 805664, 781250, 756835, 732421, 708007, 683593, 659179, 634765, 610351, 585937, 573730, 561523, 549316, 537109, 524902, 512695, 500488, 488281, 476074, 463867,
	451660, 439453, 427246, 415039, 402832, 390625, 378417, 366210, 354003, 341796, 329589, 317382, 305175, 292968, 286865, 280761, 274658, 268554, 262451, 256347, 250244, 244140, 238037, 231933, 225830,
	219726, 213623, 207519, 201416, 195312, 189208, 183105, 177001, 170898, 164794, 158691, 152587, 146484, 143432, 140380, 137329, 134277, 131225, 128173, 125122, 122070, 119018, 115966, 112915, 109863,
	106811, 103759, 100708, 97656, 94604, 91552, 88500, 85449, 82397, 79345, 76293, 73242, 71716, 70190, 68664, 67138, 65612, 64086, 62561, 61035, 59509, 57983, 56457, 54931, 53405,
	51879, 50354, 48828, 47302, 45776, 44250, 42724, 41198, 39672, 38146, 36621, 35858, 35095, 34332, 33569, 32806, 32043, 31280, 30517, 29754, 28991, 28228, 27465, 26702, 25939,
	25177, 24414, 23651, 22888, 22125, 21362, 20599, 19836, 19073, 18310, 17929, 17547, 17166, 16784, 16403, 16021, 15640, 15258, 14877, 14495, 14114, 13732, 13351, 12969, 12588,
	12207, 11825, 11444, 11062, 10681, 10299, 9918, 9536, 9155, 8964, 8773, 8583, 8392, 8201, 8010, 7820, 7629, 7438, 7247, 7057, 6866, 6675, 6484, 6294, 6103,
	5912, 5722, 5531, 5340, 5149, 4959, 4768, 4577, 4482, 4386, 4291, 4196, 4100, 4005, 3910, 3814, 3719, 3623, 3528, 3433, 3337, 3242, 3147, 3051, 2956,
	2861, 2765, 2670, 2574, 2479, 2384, 2288, 2241, 2193, 2145, 2098, 2050, 2002, 1955, 1907, 1859, 1811, 1764, 1716, 1668, 1621, 1573, 1525, 1478, 1430,
	1382, 1335, 1287, 1239, 1192, 1144, 1120, 1096, 1072, 1049, 1025, 1001, 977, 953, 929, 905, 882, 858, 834, 810, 786, 762, 739, 715, 691,
	667, 643, 619, 596, 572, 560, 548, 536, 524, 512, 500, 488, 476, 464, 452, 441, 429, 417, 405, 393, 381, 369, 357, 345, 333,
	321, 309, 298, 286, 280, 274, 268, 262, 256, 250, 244, 238, 232, 226, 220, 214, 208, 202, 196, 190, 184, 178, 172, 166, 160,
	154, 149, 143, 140, 137, 134, 131, 128, 125, 122, 119, 116, 113, 110, 107, 104, 101, 98, 95, 92, 89, 86, 83, 80, 77,
	74, 71, 70, 68, 67, 65, 64, 62, 61, 59, 58, 56, 55, 53, 52, 50, 49, 47, 46, 44, 43, 41, 40, 38, 37,
	35, 35, 34, 33, 32, 32, 31, 30, 29, 29, 28, 27, 26, 26, 25, 24, 23, 23, 22, 21, 20, 20, 19, 18, 17,
	17, 17, 16, 16, 16, 15, 15, 14, 14, 14, 13, 13, 13, 12, 12, 11, 11, 11, 10, 10, 10, 9, 9, 8, 8,
	8, 8, 8, 8, 7, 7, 7, 7, 7, 6, 6, 6, 6, 6, 5, 5, 5, 5, 5, 5, 4, 4, 4, 4, 4,
	4, 4, 4, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 0,
}

func (c *coinbaseManager) getDeflationaryPeriodBlockSubsidyFromTable(month uint64) uint64 {
	if month >= uint64(len(subsidyByDeflationaryMonthTable)) {
		month = uint64(len(subsidyByDeflationaryMonthTable) - 1)
	}
	return subsidyByDeflationaryMonthTable[month]
}

func (c *coinbaseManager) calcDeflationaryPeriodBlockSubsidyFloatCalc(month uint64) uint64 {
	baseSubsidy := c.deflationaryPhaseBaseSubsidy
	baseSubsidyCurrentPeriod := float64(baseSubsidy) / math.Pow(2, math.Trunc(float64(month)/24))
	subsidy := baseSubsidyCurrentPeriod - baseSubsidyCurrentPeriod/2/24*float64(month%24)
	return uint64(subsidy)
}

func (c *coinbaseManager) calcMergedBlockReward(stagingArea *model.StagingArea, blockHash *externalapi.DomainHash,
	blockAcceptanceData *externalapi.BlockAcceptanceData, mergingBlockDAAAddedBlocksSet hashset.HashSet) (uint64, error) {

	if !blockHash.Equal(blockAcceptanceData.BlockHash) {
		return 0, errors.Errorf("blockAcceptanceData.BlockHash is expected to be %s but got %s",
			blockHash, blockAcceptanceData.BlockHash)
	}

	if !mergingBlockDAAAddedBlocksSet.Contains(blockHash) {
		return 0, nil
	}

	totalFees := uint64(0)
	for _, txAcceptanceData := range blockAcceptanceData.TransactionAcceptanceData {
		if txAcceptanceData.IsAccepted {
			totalFees += txAcceptanceData.Fee
		}
	}

	block, err := c.blockStore.Block(c.databaseContext, stagingArea, blockHash)
	if err != nil {
		return 0, err
	}

	_, _, subsidy, err := c.ExtractCoinbaseDataBlueScoreAndSubsidy(block.Transactions[transactionhelper.CoinbaseTransactionIndex])
	if err != nil {
		return 0, err
	}

	return subsidy + totalFees, nil
}

// New instantiates a new CoinbaseManager
func New(
	databaseContext model.DBReader,

	subsidyGenesisReward uint64,
	preDeflationaryPhaseBaseSubsidy uint64,
	coinbasePayloadScriptPublicKeyMaxLength uint8,
	genesisHash *externalapi.DomainHash,
	deflationaryPhaseDaaScore uint64,
	deflationaryPhaseBaseSubsidy uint64,

	dagTraversalManager model.DAGTraversalManager,
	ghostdagDataStore model.GHOSTDAGDataStore,
	acceptanceDataStore model.AcceptanceDataStore,
	daaBlocksStore model.DAABlocksStore,
	blockStore model.BlockStore,
	pruningStore model.PruningStore,
	blockHeaderStore model.BlockHeaderStore) model.CoinbaseManager {

	return &coinbaseManager{
		databaseContext: databaseContext,

		subsidyGenesisReward:                    subsidyGenesisReward,
		preDeflationaryPhaseBaseSubsidy:         preDeflationaryPhaseBaseSubsidy,
		coinbasePayloadScriptPublicKeyMaxLength: coinbasePayloadScriptPublicKeyMaxLength,
		genesisHash:                             genesisHash,
		deflationaryPhaseDaaScore:               deflationaryPhaseDaaScore,
		deflationaryPhaseBaseSubsidy:            deflationaryPhaseBaseSubsidy,

		dagTraversalManager: dagTraversalManager,
		ghostdagDataStore:   ghostdagDataStore,
		acceptanceDataStore: acceptanceDataStore,
		daaBlocksStore:      daaBlocksStore,
		blockStore:          blockStore,
		pruningStore:        pruningStore,
		blockHeaderStore:    blockHeaderStore,
	}
}
