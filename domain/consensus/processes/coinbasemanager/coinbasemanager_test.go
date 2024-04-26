package coinbasemanager

import (
	"math"
	"strconv"
	"testing"

	"github.com/spectre-project/spectred/domain/consensus/model/externalapi"
	"github.com/spectre-project/spectred/domain/consensus/utils/constants"
	"github.com/spectre-project/spectred/domain/dagconfig"
)

func TestCalcDeflationaryPeriodBlockSubsidy(t *testing.T) {
	const secondsPerMonth = 2629800
	const secondsPerHalving = secondsPerMonth * 24
	const deflationaryPhaseDaaScore = secondsPerMonth * 6
	const deflationaryPhaseBaseSubsidy = 12 * constants.SompiPerSpectre
	coinbaseManagerInterface := New(
		nil,
		0,
		0,
		0,
		&externalapi.DomainHash{},
		deflationaryPhaseDaaScore,
		deflationaryPhaseBaseSubsidy,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil)
	coinbaseManagerInstance := coinbaseManagerInterface.(*coinbaseManager)

	tests := []struct {
		name                 string
		blockDaaScore        uint64
		expectedBlockSubsidy uint64
	}{
		{
			name:                 "start of deflationary phase",
			blockDaaScore:        deflationaryPhaseDaaScore,
			expectedBlockSubsidy: deflationaryPhaseBaseSubsidy,
		},
		{
			name:                 "after 2 years",
			blockDaaScore:        deflationaryPhaseDaaScore + secondsPerHalving,
			expectedBlockSubsidy: uint64(math.Trunc(deflationaryPhaseBaseSubsidy / 2)),
		},
		{
			name:                 "after 4 years",
			blockDaaScore:        deflationaryPhaseDaaScore + secondsPerHalving*2,
			expectedBlockSubsidy: uint64(math.Trunc(deflationaryPhaseBaseSubsidy / 4)),
		},
		{
			name:                 "after 8 years",
			blockDaaScore:        deflationaryPhaseDaaScore + secondsPerHalving*4,
			expectedBlockSubsidy: uint64(math.Trunc(deflationaryPhaseBaseSubsidy / 16)),
		},
		{
			name:                 "after 16 years",
			blockDaaScore:        deflationaryPhaseDaaScore + secondsPerHalving*8,
			expectedBlockSubsidy: uint64(math.Trunc(deflationaryPhaseBaseSubsidy / 256)),
		},
		{
			name:                 "after 32 years",
			blockDaaScore:        deflationaryPhaseDaaScore + secondsPerHalving*16,
			expectedBlockSubsidy: uint64(math.Trunc(deflationaryPhaseBaseSubsidy / 65536)),
		},
		{
			name:                 "just before subsidy depleted",
			blockDaaScore:        deflationaryPhaseDaaScore + (secondsPerHalving / 24 * 725),
			expectedBlockSubsidy: 1,
		},
		{
			name:                 "after subsidy depleted",
			blockDaaScore:        deflationaryPhaseDaaScore + (secondsPerHalving / 24 * 726),
			expectedBlockSubsidy: 0,
		},
	}

	for _, test := range tests {
		blockSubsidy := coinbaseManagerInstance.calcDeflationaryPeriodBlockSubsidy(test.blockDaaScore)
		if blockSubsidy != test.expectedBlockSubsidy {
			t.Errorf("TestCalcDeflationaryPeriodBlockSubsidy: test '%s' failed. Want: %d, got: %d",
				test.name, test.expectedBlockSubsidy, blockSubsidy)
		}
	}
}

func TestBuildSubsidyTable(t *testing.T) {
	deflationaryPhaseBaseSubsidy := dagconfig.MainnetParams.DeflationaryPhaseBaseSubsidy
	if deflationaryPhaseBaseSubsidy != 12*constants.SompiPerSpectre {
		t.Errorf("TestBuildSubsidyTable: table generation function was not updated to reflect "+
			"the new base subsidy %d. Please fix the constant above and replace subsidyByDeflationaryMonthTable "+
			"in coinbasemanager.go with the printed table", deflationaryPhaseBaseSubsidy)
	}
	coinbaseManagerInterface := New(
		nil,
		0,
		0,
		0,
		&externalapi.DomainHash{},
		0,
		deflationaryPhaseBaseSubsidy,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil)
	coinbaseManagerInstance := coinbaseManagerInterface.(*coinbaseManager)

	var subsidyTable []uint64
	for M := uint64(0); ; M++ {
		subsidy := coinbaseManagerInstance.calcDeflationaryPeriodBlockSubsidyFloatCalc(M)
		subsidyTable = append(subsidyTable, subsidy)
		if subsidy == 0 {
			break
		}
	}

	tableStr := "\n{\t"
	for i := 0; i < len(subsidyTable); i++ {
		tableStr += strconv.FormatUint(subsidyTable[i], 10) + ", "
		if (i+1)%25 == 0 {
			tableStr += "\n\t"
		}
	}
	tableStr += "\n}"
	t.Logf(tableStr)
}
