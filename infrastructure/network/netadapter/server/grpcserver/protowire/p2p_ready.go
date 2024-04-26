package protowire

import (
	"github.com/pkg/errors"
	"github.com/spectre-project/spectred/app/appmessage"
)

func (x *SpectredMessage_Ready) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "SpectredMessage_Ready is nil")
	}
	return &appmessage.MsgReady{}, nil
}

func (x *SpectredMessage_Ready) fromAppMessage(_ *appmessage.MsgReady) error {
	return nil
}
