package protowire

import (
	"github.com/pkg/errors"
	"github.com/spectre-project/spectred/app/appmessage"
)

type converter interface {
	toAppMessage() (appmessage.Message, error)
}

// ToAppMessage converts a SpectredMessage to its appmessage.Message representation
func (x *SpectredMessage) ToAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "SpectredMessage is nil")
	}
	converter, ok := x.Payload.(converter)
	if !ok {
		return nil, errors.Errorf("received invalid message")
	}
	appMessage, err := converter.toAppMessage()
	if err != nil {
		return nil, err
	}
	return appMessage, nil
}

// FromAppMessage creates a SpectredMessage from a appmessage.Message
func FromAppMessage(message appmessage.Message) (*SpectredMessage, error) {
	payload, err := toPayload(message)
	if err != nil {
		return nil, err
	}
	return &SpectredMessage{
		Payload: payload,
	}, nil
}

func toPayload(message appmessage.Message) (isSpectredMessage_Payload, error) {
	p2pPayload, err := toP2PPayload(message)
	if err != nil {
		return nil, err
	}
	if p2pPayload != nil {
		return p2pPayload, nil
	}

	rpcPayload, err := toRPCPayload(message)
	if err != nil {
		return nil, err
	}
	if rpcPayload != nil {
		return rpcPayload, nil
	}

	return nil, errors.Errorf("unknown message type %T", message)
}

func toP2PPayload(message appmessage.Message) (isSpectredMessage_Payload, error) {
	switch message := message.(type) {
	case *appmessage.MsgAddresses:
		payload := new(SpectredMessage_Addresses)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgBlock:
		payload := new(SpectredMessage_Block)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestBlockLocator:
		payload := new(SpectredMessage_RequestBlockLocator)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgBlockLocator:
		payload := new(SpectredMessage_BlockLocator)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestAddresses:
		payload := new(SpectredMessage_RequestAddresses)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestIBDBlocks:
		payload := new(SpectredMessage_RequestIBDBlocks)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestNextHeaders:
		payload := new(SpectredMessage_RequestNextHeaders)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgDoneHeaders:
		payload := new(SpectredMessage_DoneHeaders)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestRelayBlocks:
		payload := new(SpectredMessage_RequestRelayBlocks)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestTransactions:
		payload := new(SpectredMessage_RequestTransactions)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgTransactionNotFound:
		payload := new(SpectredMessage_TransactionNotFound)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgInvRelayBlock:
		payload := new(SpectredMessage_InvRelayBlock)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgInvTransaction:
		payload := new(SpectredMessage_InvTransactions)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgPing:
		payload := new(SpectredMessage_Ping)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgPong:
		payload := new(SpectredMessage_Pong)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgTx:
		payload := new(SpectredMessage_Transaction)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgVerAck:
		payload := new(SpectredMessage_Verack)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgVersion:
		payload := new(SpectredMessage_Version)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgReject:
		payload := new(SpectredMessage_Reject)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestPruningPointUTXOSet:
		payload := new(SpectredMessage_RequestPruningPointUTXOSet)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgPruningPointUTXOSetChunk:
		payload := new(SpectredMessage_PruningPointUtxoSetChunk)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgUnexpectedPruningPoint:
		payload := new(SpectredMessage_UnexpectedPruningPoint)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgIBDBlockLocator:
		payload := new(SpectredMessage_IbdBlockLocator)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgIBDBlockLocatorHighestHash:
		payload := new(SpectredMessage_IbdBlockLocatorHighestHash)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgIBDBlockLocatorHighestHashNotFound:
		payload := new(SpectredMessage_IbdBlockLocatorHighestHashNotFound)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.BlockHeadersMessage:
		payload := new(SpectredMessage_BlockHeaders)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestNextPruningPointUTXOSetChunk:
		payload := new(SpectredMessage_RequestNextPruningPointUtxoSetChunk)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgDonePruningPointUTXOSetChunks:
		payload := new(SpectredMessage_DonePruningPointUtxoSetChunks)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgBlockWithTrustedData:
		payload := new(SpectredMessage_BlockWithTrustedData)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestPruningPointAndItsAnticone:
		payload := new(SpectredMessage_RequestPruningPointAndItsAnticone)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgDoneBlocksWithTrustedData:
		payload := new(SpectredMessage_DoneBlocksWithTrustedData)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgIBDBlock:
		payload := new(SpectredMessage_IbdBlock)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestHeaders:
		payload := new(SpectredMessage_RequestHeaders)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgPruningPoints:
		payload := new(SpectredMessage_PruningPoints)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestPruningPointProof:
		payload := new(SpectredMessage_RequestPruningPointProof)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgPruningPointProof:
		payload := new(SpectredMessage_PruningPointProof)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgReady:
		payload := new(SpectredMessage_Ready)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgTrustedData:
		payload := new(SpectredMessage_TrustedData)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgBlockWithTrustedDataV4:
		payload := new(SpectredMessage_BlockWithTrustedDataV4)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestNextPruningPointAndItsAnticoneBlocks:
		payload := new(SpectredMessage_RequestNextPruningPointAndItsAnticoneBlocks)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestIBDChainBlockLocator:
		payload := new(SpectredMessage_RequestIBDChainBlockLocator)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgIBDChainBlockLocator:
		payload := new(SpectredMessage_IbdChainBlockLocator)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestAnticone:
		payload := new(SpectredMessage_RequestAnticone)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	default:
		return nil, nil
	}
}

func toRPCPayload(message appmessage.Message) (isSpectredMessage_Payload, error) {
	switch message := message.(type) {
	case *appmessage.GetCurrentNetworkRequestMessage:
		payload := new(SpectredMessage_GetCurrentNetworkRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetCurrentNetworkResponseMessage:
		payload := new(SpectredMessage_GetCurrentNetworkResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.SubmitBlockRequestMessage:
		payload := new(SpectredMessage_SubmitBlockRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.SubmitBlockResponseMessage:
		payload := new(SpectredMessage_SubmitBlockResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBlockTemplateRequestMessage:
		payload := new(SpectredMessage_GetBlockTemplateRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBlockTemplateResponseMessage:
		payload := new(SpectredMessage_GetBlockTemplateResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyBlockAddedRequestMessage:
		payload := new(SpectredMessage_NotifyBlockAddedRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyBlockAddedResponseMessage:
		payload := new(SpectredMessage_NotifyBlockAddedResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.BlockAddedNotificationMessage:
		payload := new(SpectredMessage_BlockAddedNotification)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetPeerAddressesRequestMessage:
		payload := new(SpectredMessage_GetPeerAddressesRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetPeerAddressesResponseMessage:
		payload := new(SpectredMessage_GetPeerAddressesResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetSelectedTipHashRequestMessage:
		payload := new(SpectredMessage_GetSelectedTipHashRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetSelectedTipHashResponseMessage:
		payload := new(SpectredMessage_GetSelectedTipHashResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetMempoolEntryRequestMessage:
		payload := new(SpectredMessage_GetMempoolEntryRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetMempoolEntryResponseMessage:
		payload := new(SpectredMessage_GetMempoolEntryResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetConnectedPeerInfoRequestMessage:
		payload := new(SpectredMessage_GetConnectedPeerInfoRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetConnectedPeerInfoResponseMessage:
		payload := new(SpectredMessage_GetConnectedPeerInfoResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.AddPeerRequestMessage:
		payload := new(SpectredMessage_AddPeerRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.AddPeerResponseMessage:
		payload := new(SpectredMessage_AddPeerResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.SubmitTransactionRequestMessage:
		payload := new(SpectredMessage_SubmitTransactionRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.SubmitTransactionResponseMessage:
		payload := new(SpectredMessage_SubmitTransactionResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyVirtualSelectedParentChainChangedRequestMessage:
		payload := new(SpectredMessage_NotifyVirtualSelectedParentChainChangedRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyVirtualSelectedParentChainChangedResponseMessage:
		payload := new(SpectredMessage_NotifyVirtualSelectedParentChainChangedResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.VirtualSelectedParentChainChangedNotificationMessage:
		payload := new(SpectredMessage_VirtualSelectedParentChainChangedNotification)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBlockRequestMessage:
		payload := new(SpectredMessage_GetBlockRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBlockResponseMessage:
		payload := new(SpectredMessage_GetBlockResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetSubnetworkRequestMessage:
		payload := new(SpectredMessage_GetSubnetworkRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetSubnetworkResponseMessage:
		payload := new(SpectredMessage_GetSubnetworkResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetVirtualSelectedParentChainFromBlockRequestMessage:
		payload := new(SpectredMessage_GetVirtualSelectedParentChainFromBlockRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetVirtualSelectedParentChainFromBlockResponseMessage:
		payload := new(SpectredMessage_GetVirtualSelectedParentChainFromBlockResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBlocksRequestMessage:
		payload := new(SpectredMessage_GetBlocksRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBlocksResponseMessage:
		payload := new(SpectredMessage_GetBlocksResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBlockCountRequestMessage:
		payload := new(SpectredMessage_GetBlockCountRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBlockCountResponseMessage:
		payload := new(SpectredMessage_GetBlockCountResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBlockDAGInfoRequestMessage:
		payload := new(SpectredMessage_GetBlockDagInfoRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBlockDAGInfoResponseMessage:
		payload := new(SpectredMessage_GetBlockDagInfoResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.ResolveFinalityConflictRequestMessage:
		payload := new(SpectredMessage_ResolveFinalityConflictRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.ResolveFinalityConflictResponseMessage:
		payload := new(SpectredMessage_ResolveFinalityConflictResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyFinalityConflictsRequestMessage:
		payload := new(SpectredMessage_NotifyFinalityConflictsRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyFinalityConflictsResponseMessage:
		payload := new(SpectredMessage_NotifyFinalityConflictsResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.FinalityConflictNotificationMessage:
		payload := new(SpectredMessage_FinalityConflictNotification)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.FinalityConflictResolvedNotificationMessage:
		payload := new(SpectredMessage_FinalityConflictResolvedNotification)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetMempoolEntriesRequestMessage:
		payload := new(SpectredMessage_GetMempoolEntriesRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetMempoolEntriesResponseMessage:
		payload := new(SpectredMessage_GetMempoolEntriesResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.ShutDownRequestMessage:
		payload := new(SpectredMessage_ShutDownRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.ShutDownResponseMessage:
		payload := new(SpectredMessage_ShutDownResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetHeadersRequestMessage:
		payload := new(SpectredMessage_GetHeadersRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetHeadersResponseMessage:
		payload := new(SpectredMessage_GetHeadersResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyUTXOsChangedRequestMessage:
		payload := new(SpectredMessage_NotifyUtxosChangedRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyUTXOsChangedResponseMessage:
		payload := new(SpectredMessage_NotifyUtxosChangedResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.UTXOsChangedNotificationMessage:
		payload := new(SpectredMessage_UtxosChangedNotification)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.StopNotifyingUTXOsChangedRequestMessage:
		payload := new(SpectredMessage_StopNotifyingUtxosChangedRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.StopNotifyingUTXOsChangedResponseMessage:
		payload := new(SpectredMessage_StopNotifyingUtxosChangedResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetUTXOsByAddressesRequestMessage:
		payload := new(SpectredMessage_GetUtxosByAddressesRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetUTXOsByAddressesResponseMessage:
		payload := new(SpectredMessage_GetUtxosByAddressesResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBalanceByAddressRequestMessage:
		payload := new(SpectredMessage_GetBalanceByAddressRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBalanceByAddressResponseMessage:
		payload := new(SpectredMessage_GetBalanceByAddressResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetVirtualSelectedParentBlueScoreRequestMessage:
		payload := new(SpectredMessage_GetVirtualSelectedParentBlueScoreRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetVirtualSelectedParentBlueScoreResponseMessage:
		payload := new(SpectredMessage_GetVirtualSelectedParentBlueScoreResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyVirtualSelectedParentBlueScoreChangedRequestMessage:
		payload := new(SpectredMessage_NotifyVirtualSelectedParentBlueScoreChangedRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyVirtualSelectedParentBlueScoreChangedResponseMessage:
		payload := new(SpectredMessage_NotifyVirtualSelectedParentBlueScoreChangedResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.VirtualSelectedParentBlueScoreChangedNotificationMessage:
		payload := new(SpectredMessage_VirtualSelectedParentBlueScoreChangedNotification)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.BanRequestMessage:
		payload := new(SpectredMessage_BanRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.BanResponseMessage:
		payload := new(SpectredMessage_BanResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.UnbanRequestMessage:
		payload := new(SpectredMessage_UnbanRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.UnbanResponseMessage:
		payload := new(SpectredMessage_UnbanResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetInfoRequestMessage:
		payload := new(SpectredMessage_GetInfoRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetInfoResponseMessage:
		payload := new(SpectredMessage_GetInfoResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyPruningPointUTXOSetOverrideRequestMessage:
		payload := new(SpectredMessage_NotifyPruningPointUTXOSetOverrideRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyPruningPointUTXOSetOverrideResponseMessage:
		payload := new(SpectredMessage_NotifyPruningPointUTXOSetOverrideResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.PruningPointUTXOSetOverrideNotificationMessage:
		payload := new(SpectredMessage_PruningPointUTXOSetOverrideNotification)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.StopNotifyingPruningPointUTXOSetOverrideRequestMessage:
		payload := new(SpectredMessage_StopNotifyingPruningPointUTXOSetOverrideRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.EstimateNetworkHashesPerSecondRequestMessage:
		payload := new(SpectredMessage_EstimateNetworkHashesPerSecondRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.EstimateNetworkHashesPerSecondResponseMessage:
		payload := new(SpectredMessage_EstimateNetworkHashesPerSecondResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyVirtualDaaScoreChangedRequestMessage:
		payload := new(SpectredMessage_NotifyVirtualDaaScoreChangedRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyVirtualDaaScoreChangedResponseMessage:
		payload := new(SpectredMessage_NotifyVirtualDaaScoreChangedResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.VirtualDaaScoreChangedNotificationMessage:
		payload := new(SpectredMessage_VirtualDaaScoreChangedNotification)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBalancesByAddressesRequestMessage:
		payload := new(SpectredMessage_GetBalancesByAddressesRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBalancesByAddressesResponseMessage:
		payload := new(SpectredMessage_GetBalancesByAddressesResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyNewBlockTemplateRequestMessage:
		payload := new(SpectredMessage_NotifyNewBlockTemplateRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyNewBlockTemplateResponseMessage:
		payload := new(SpectredMessage_NotifyNewBlockTemplateResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NewBlockTemplateNotificationMessage:
		payload := new(SpectredMessage_NewBlockTemplateNotification)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetMempoolEntriesByAddressesRequestMessage:
		payload := new(SpectredMessage_GetMempoolEntriesByAddressesRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetMempoolEntriesByAddressesResponseMessage:
		payload := new(SpectredMessage_GetMempoolEntriesByAddressesResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetCoinSupplyRequestMessage:
		payload := new(SpectredMessage_GetCoinSupplyRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetCoinSupplyResponseMessage:
		payload := new(SpectredMessage_GetCoinSupplyResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	default:
		return nil, nil
	}
}
