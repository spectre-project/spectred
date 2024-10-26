package protowire

import (
	"github.com/pkg/errors"
	"github.com/spectre-project/spectred/app/appmessage"
)

func (x *SpectredMessage_SubmitTransactionReplacementRequest) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "SpectredMessage_SubmitTransactionReplacementRequest is nil")
	}
	return x.SubmitTransactionReplacementRequest.toAppMessage()
}

func (x *SpectredMessage_SubmitTransactionReplacementRequest) fromAppMessage(message *appmessage.SubmitTransactionReplacementRequestMessage) error {
	x.SubmitTransactionReplacementRequest = &SubmitTransactionReplacementRequestMessage{
		Transaction: &RpcTransaction{},
	}
	x.SubmitTransactionReplacementRequest.Transaction.fromAppMessage(message.Transaction)
	return nil
}

func (x *SubmitTransactionReplacementRequestMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "SubmitBlockRequestMessage is nil")
	}
	rpcTransaction, err := x.Transaction.toAppMessage()
	if err != nil {
		return nil, err
	}
	return &appmessage.SubmitTransactionReplacementRequestMessage{
		Transaction: rpcTransaction,
	}, nil
}

func (x *SpectredMessage_SubmitTransactionReplacementResponse) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "SpectredMessage_SubmitTransactionReplacementResponse is nil")
	}
	return x.SubmitTransactionReplacementResponse.toAppMessage()
}

func (x *SpectredMessage_SubmitTransactionReplacementResponse) fromAppMessage(message *appmessage.SubmitTransactionReplacementResponseMessage) error {
	var err *RPCError
	if message.Error != nil {
		err = &RPCError{Message: message.Error.Message}
	}
	x.SubmitTransactionReplacementResponse = &SubmitTransactionReplacementResponseMessage{
		TransactionId:       message.TransactionID,
		ReplacedTransaction: &RpcTransaction{},
		Error:               err,
	}
	if message.ReplacedTransaction != nil {
		x.SubmitTransactionReplacementResponse.ReplacedTransaction.fromAppMessage(message.ReplacedTransaction)
	}
	return nil
}

func (x *SubmitTransactionReplacementResponseMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "SubmitTransactionReplacementResponseMessage is nil")
	}

	if x.Error != nil {
		rpcErr, err := x.Error.toAppMessage()
		// Error is an optional field
		if err != nil && !errors.Is(err, errorNil) {
			return nil, err
		}

		return &appmessage.SubmitTransactionReplacementResponseMessage{
			TransactionID: x.TransactionId,
			Error:         rpcErr,
		}, nil
	}

	replacedTx, err := x.ReplacedTransaction.toAppMessage()
	if err != nil {
		return nil, err
	}
	return &appmessage.SubmitTransactionReplacementResponseMessage{
		TransactionID:       x.TransactionId,
		ReplacedTransaction: replacedTx,
	}, nil
}
