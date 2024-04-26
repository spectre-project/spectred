package testapi

import (
	"github.com/spectre-project/spectred/domain/consensus/model"
	"github.com/spectre-project/spectred/domain/consensus/utils/txscript"
)

// TestTransactionValidator adds to the main TransactionValidator methods required by tests
type TestTransactionValidator interface {
	model.TransactionValidator
	SigCache() *txscript.SigCache
	SetSigCache(sigCache *txscript.SigCache)
}
