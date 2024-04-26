package protowire

import (
	"github.com/pkg/errors"
	"github.com/spectre-project/spectred/app/appmessage"
)

func (x *SpectredMessage_IbdBlock) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "SpectredMessage_IbdBlock is nil")
	}
	msgBlock, err := x.IbdBlock.toAppMessage()
	if err != nil {
		return nil, err
	}
	return &appmessage.MsgIBDBlock{MsgBlock: msgBlock}, nil
}

func (x *SpectredMessage_IbdBlock) fromAppMessage(msgIBDBlock *appmessage.MsgIBDBlock) error {
	x.IbdBlock = new(BlockMessage)
	return x.IbdBlock.fromAppMessage(msgIBDBlock.MsgBlock)
}
