package protowire

import (
	"github.com/pkg/errors"
	"github.com/spectre-project/spectred/app/appmessage"
)

func (x *SpectredMessage_GetHeadersRequest) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "SpectredMessage_GetHeadersRequest is nil")
	}
	return x.GetHeadersRequest.toAppMessage()
}

func (x *SpectredMessage_GetHeadersRequest) fromAppMessage(message *appmessage.GetHeadersRequestMessage) error {
	x.GetHeadersRequest = &GetHeadersRequestMessage{
		StartHash:   message.StartHash,
		Limit:       message.Limit,
		IsAscending: message.IsAscending,
	}
	return nil
}

func (x *GetHeadersRequestMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "GetHeadersRequestMessage is nil")
	}
	return &appmessage.GetHeadersRequestMessage{
		StartHash:   x.StartHash,
		Limit:       x.Limit,
		IsAscending: x.IsAscending,
	}, nil
}

func (x *SpectredMessage_GetHeadersResponse) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "SpectredMessage_GetHeadersResponse is nil")
	}
	return x.GetHeadersResponse.toAppMessage()
}

func (x *SpectredMessage_GetHeadersResponse) fromAppMessage(message *appmessage.GetHeadersResponseMessage) error {
	var err *RPCError
	if message.Error != nil {
		err = &RPCError{Message: message.Error.Message}
	}
	x.GetHeadersResponse = &GetHeadersResponseMessage{
		Headers: message.Headers,
		Error:   err,
	}
	return nil
}

func (x *GetHeadersResponseMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "GetHeadersResponseMessage is nil")
	}
	rpcErr, err := x.Error.toAppMessage()
	// Error is an optional field
	if err != nil && !errors.Is(err, errorNil) {
		return nil, err
	}

	if rpcErr != nil && len(x.Headers) != 0 {
		return nil, errors.New("GetHeadersResponseMessage contains both an error and a response")
	}

	return &appmessage.GetHeadersResponseMessage{
		Headers: x.Headers,
		Error:   rpcErr,
	}, nil
}
