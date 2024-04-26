package protowire

import (
	"github.com/pkg/errors"
	"github.com/spectre-project/spectred/app/appmessage"
)

func (x *SpectredMessage_NotifyNewBlockTemplateRequest) toAppMessage() (appmessage.Message, error) {
	return &appmessage.NotifyNewBlockTemplateRequestMessage{}, nil
}

func (x *SpectredMessage_NotifyNewBlockTemplateRequest) fromAppMessage(_ *appmessage.NotifyNewBlockTemplateRequestMessage) error {
	x.NotifyNewBlockTemplateRequest = &NotifyNewBlockTemplateRequestMessage{}
	return nil
}

func (x *SpectredMessage_NotifyNewBlockTemplateResponse) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "SpectredMessage_NotifyNewBlockTemplateResponse is nil")
	}
	return x.NotifyNewBlockTemplateResponse.toAppMessage()
}

func (x *SpectredMessage_NotifyNewBlockTemplateResponse) fromAppMessage(message *appmessage.NotifyNewBlockTemplateResponseMessage) error {
	var err *RPCError
	if message.Error != nil {
		err = &RPCError{Message: message.Error.Message}
	}
	x.NotifyNewBlockTemplateResponse = &NotifyNewBlockTemplateResponseMessage{
		Error: err,
	}
	return nil
}

func (x *NotifyNewBlockTemplateResponseMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "NotifyNewBlockTemplateResponseMessage is nil")
	}
	rpcErr, err := x.Error.toAppMessage()
	// Error is an optional field
	if err != nil && !errors.Is(err, errorNil) {
		return nil, err
	}
	return &appmessage.NotifyNewBlockTemplateResponseMessage{
		Error: rpcErr,
	}, nil
}

func (x *SpectredMessage_NewBlockTemplateNotification) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "SpectredMessage_NewBlockTemplateNotification is nil")
	}
	return x.NewBlockTemplateNotification.toAppMessage()
}

func (x *SpectredMessage_NewBlockTemplateNotification) fromAppMessage(message *appmessage.NewBlockTemplateNotificationMessage) error {
	x.NewBlockTemplateNotification = &NewBlockTemplateNotificationMessage{}
	return nil
}

func (x *NewBlockTemplateNotificationMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "NewBlockTemplateNotificationMessage is nil")
	}
	return &appmessage.NewBlockTemplateNotificationMessage{}, nil
}
