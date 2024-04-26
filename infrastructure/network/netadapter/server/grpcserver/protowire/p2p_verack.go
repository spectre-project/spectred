package protowire

import (
	"github.com/pkg/errors"
	"github.com/spectre-project/spectred/app/appmessage"
)

func (x *SpectredMessage_Verack) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "SpectredMessage_Verack is nil")
	}
	return &appmessage.MsgVerAck{}, nil
}

func (x *SpectredMessage_Verack) fromAppMessage(_ *appmessage.MsgVerAck) error {
	return nil
}
