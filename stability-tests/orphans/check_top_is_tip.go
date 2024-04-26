package main

import (
	"github.com/pkg/errors"
	"github.com/spectre-project/spectred/domain/consensus/model/externalapi"
	"github.com/spectre-project/spectred/domain/consensus/utils/consensushashing"
	"github.com/spectre-project/spectred/stability-tests/common/rpc"
)

func checkTopBlockIsTip(rpcClient *rpc.Client, topBlock *externalapi.DomainBlock) error {
	selectedTipHashResponse, err := rpcClient.GetSelectedTipHash()
	if err != nil {
		return err
	}

	topBlockHash := consensushashing.BlockHash(topBlock)
	if selectedTipHashResponse.SelectedTipHash != topBlockHash.String() {
		return errors.Errorf("selectedTipHash is '%s' while expected to be topBlock's hash `%s`",
			selectedTipHashResponse.SelectedTipHash, topBlockHash)
	}

	return nil
}
