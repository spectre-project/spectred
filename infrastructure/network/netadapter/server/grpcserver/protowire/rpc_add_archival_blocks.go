package protowire

import (
	"github.com/pkg/errors"
	"github.com/spectre-project/spectred/app/appmessage"
)

func (x *SpectredMessage_AddArchivalBlocksRequest) toAppMessage() (appmessage.Message, error) {
	panic("we need to implement acceptance data conversion")
	if x == nil {
		return nil, errors.Wrapf(errorNil, "SpectredMessage_AddArchivalBlocksRequest is nil")
	}

	blocks := make([]*appmessage.ArchivalBlock, len(x.AddArchivalBlocksRequest.Blocks))
	for i, block := range x.AddArchivalBlocksRequest.Blocks {
		rpcBlock, err := block.Block.toAppMessage()
		if err != nil {
			return nil, err
		}

		blocks[i] = &appmessage.ArchivalBlock{
			Block: rpcBlock,
			Child: block.Child,
		}
	}

	return &appmessage.AddArchivalBlocksRequestMessage{
		Blocks: blocks,
	}, nil
}

func (x *SpectredMessage_AddArchivalBlocksRequest) fromAppMessage(message *appmessage.AddArchivalBlocksRequestMessage) error {
	blocks := make([]*ArchivalBlock, len(message.Blocks))
	for i, block := range message.Blocks {
		protoBlock := &ArchivalBlock{
			Child:          block.Child,
			SelectedParent: block.SelectedParent,
		}

		if block.Block != nil {
			protoBlock.Block = &RpcBlock{}
			err := protoBlock.Block.fromAppMessage(block.Block)
			if err != nil {
				return err
			}
		}

		protoBlock.AcceptanceData = make([]*MergesetBlockAcceptanceData, len(block.AcceptanceData))
		for j, acceptanceData := range block.AcceptanceData {
			protoBlock.AcceptanceData[j] = &MergesetBlockAcceptanceData{}
			protoBlock.AcceptanceData[j].fromAppMessage(acceptanceData)
		}

		blocks[i] = protoBlock
	}

	x.AddArchivalBlocksRequest = &AddArchivalBlocksRequestMessage{
		Blocks: blocks,
	}
	return nil
}

func (x *MergesetBlockAcceptanceData) fromAppMessage(message *appmessage.MergesetBlockAcceptanceData) error {
	if message == nil {
		return errors.Wrapf(errorNil, "MergesetBlockAcceptanceData is nil")
	}

	x.BlockHash = message.BlockHash
	x.AcceptedTxs = make([]*AcceptedTxEntry, len(message.AcceptedTxs))
	for i, tx := range message.AcceptedTxs {
		x.AcceptedTxs[i] = &AcceptedTxEntry{
			TransactionId:    tx.TransactionID,
			IndexWithinBlock: tx.IndexWithinBlock,
		}
	}

	return nil
}

func (x *SpectredMessage_AddArchivalBlocksResponse) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "SpectredMessage_AddArchivalBlocksResponse is nil")
	}
	return x.AddArchivalBlocksResponse.toAppMessage()
}

func (x *SpectredMessage_AddArchivalBlocksResponse) fromAppMessage(message *appmessage.AddArchivalBlocksResponseMessage) error {
	var err *RPCError
	if message.Error != nil {
		err = &RPCError{Message: message.Error.Message}
	}

	x.AddArchivalBlocksResponse = &AddArchivalBlocksResponseMessage{
		Error: err,
	}

	return nil
}

func (x *AddArchivalBlocksResponseMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "AddArchivalBlocksResponseMessage is nil")
	}
	rpcErr, err := x.Error.toAppMessage()
	// Error is an optional field
	if err != nil && !errors.Is(err, errorNil) {
		return nil, err
	}

	return &appmessage.AddArchivalBlocksResponseMessage{
		Error: rpcErr,
	}, nil
}
