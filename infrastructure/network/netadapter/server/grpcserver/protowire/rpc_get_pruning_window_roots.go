package protowire

import (
	"github.com/pkg/errors"
	"github.com/spectre-project/spectred/app/appmessage"
)

func (x *SpectredMessage_GetPruningWindowRootsRequest) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "SpectredMessage_GetPruningWindowRootsRequest is nil")
	}
	return &appmessage.GetPeerAddressesRequestMessage{}, nil
}

func (x *SpectredMessage_GetPruningWindowRootsRequest) fromAppMessage(_ *appmessage.GetPruningWindowRootsRequestMessage) error {
	return nil
}

func (x *SpectredMessage_GetPruningWindowRootsResponse) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "SpectredMessage_GetPruningWindowRootsResponse is nil")
	}
	return x.GetPruningWindowRootsResponse.toAppMessage()
}

func (x *SpectredMessage_GetPruningWindowRootsResponse) fromAppMessage(message *appmessage.GetPruningWindowRootsResponseMessage) error {
	var err *RPCError
	if message.Error != nil {
		err = &RPCError{Message: message.Error.Message}
	}

	roots := make([]*PruningWindowRoots, len(message.Roots))
	for i, root := range message.Roots {
		roots[i] = &PruningWindowRoots{
			PpRoots: root.PPRoots,
			PpIndex: root.PPIndex,
		}
	}

	x.GetPruningWindowRootsResponse = &GetPruningWindowRootsResponseMessage{
		Roots: roots,
		Error: err,
	}

	return nil
}

func (x *GetPruningWindowRootsResponseMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "GetPeerAddressesResponseMessage is nil")
	}
	rpcErr, err := x.Error.toAppMessage()
	// Error is an optional field
	if err != nil && !errors.Is(err, errorNil) {
		return nil, err
	}

	roots := make([]*appmessage.PruningWindowRoots, len(x.Roots))
	for i, root := range x.Roots {
		roots[i] = &appmessage.PruningWindowRoots{
			PPRoots: root.PpRoots,
			PPIndex: root.PpIndex,
		}
	}

	return &appmessage.GetPruningWindowRootsResponseMessage{
		Roots: roots,
		Error: rpcErr,
	}, nil
}
