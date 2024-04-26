package rpchandlers

import (
	"github.com/spectre-project/spectred/app/appmessage"
	"github.com/spectre-project/spectred/app/rpc/rpccontext"
	"github.com/spectre-project/spectred/infrastructure/network/netadapter/router"
)

// HandleNotifyVirtualSelectedParentBlueScoreChanged handles the respectively named RPC command
func HandleNotifyVirtualSelectedParentBlueScoreChanged(context *rpccontext.Context, router *router.Router, _ appmessage.Message) (appmessage.Message, error) {
	listener, err := context.NotificationManager.Listener(router)
	if err != nil {
		return nil, err
	}
	listener.PropagateVirtualSelectedParentBlueScoreChangedNotifications()

	response := appmessage.NewNotifyVirtualSelectedParentBlueScoreChangedResponseMessage()
	return response, nil
}
