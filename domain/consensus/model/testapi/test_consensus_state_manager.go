package testapi

import (
	"github.com/spectre-project/spectred/domain/consensus/model"
	"github.com/spectre-project/spectred/domain/consensus/model/externalapi"
)

// TestConsensusStateManager  adds to the main ConsensusStateManager methods required by tests
type TestConsensusStateManager interface {
	model.ConsensusStateManager
	AddUTXOToMultiset(multiset model.Multiset, entry externalapi.UTXOEntry,
		outpoint *externalapi.DomainOutpoint) error
	ResolveBlockStatus(stagingArea *model.StagingArea, blockHash *externalapi.DomainHash,
		useSeparateStagingAreaPerBlock bool) (externalapi.BlockStatus, error)
}
