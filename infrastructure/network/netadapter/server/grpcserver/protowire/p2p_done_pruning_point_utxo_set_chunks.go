package protowire

import (
	"github.com/pkg/errors"
	"github.com/spectre-project/spectred/app/appmessage"
)

func (x *SpectredMessage_DonePruningPointUtxoSetChunks) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "SpectredMessage_DonePruningPointUtxoSetChunks is nil")
	}
	return &appmessage.MsgDonePruningPointUTXOSetChunks{}, nil
}

func (x *SpectredMessage_DonePruningPointUtxoSetChunks) fromAppMessage(_ *appmessage.MsgDonePruningPointUTXOSetChunks) error {
	x.DonePruningPointUtxoSetChunks = &DonePruningPointUtxoSetChunksMessage{}
	return nil
}
