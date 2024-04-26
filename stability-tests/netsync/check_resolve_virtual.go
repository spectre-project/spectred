package main

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/spectre-project/spectred/app/appmessage"
	"github.com/spectre-project/spectred/stability-tests/common/rpc"
)

func checkResolveVirtual(syncerClient, syncedClient *rpc.Client) error {
	err := syncedClient.RegisterForBlockAddedNotifications()
	if err != nil {
		return errors.Wrap(err, "error registering for blockAdded notifications")
	}

	syncedBlockCountResponse, err := syncedClient.GetBlockCount()
	if err != nil {
		return err
	}

	rejectReason, err := mineOnTips(syncerClient)
	if err != nil {
		panic(err)
	}
	if rejectReason != appmessage.RejectReasonNone {
		panic(fmt.Sprintf("mined block rejected: %s", rejectReason))
	}

	expectedDuration := time.Duration(syncedBlockCountResponse.BlockCount) * 100 * time.Millisecond
	start := time.Now()
	select {
	case <-time.After(expectedDuration):
		return errors.Errorf("it took more than %s to resolve the virtual", expectedDuration)
	case <-syncedClient.OnBlockAdded:
	}

	log.Infof("It took %s to resolve the virtual", time.Since(start))
	return nil
}
