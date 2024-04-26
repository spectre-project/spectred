package main

import (
	"context"
	"fmt"

	"github.com/spectre-project/spectred/cmd/spectrewallet/daemon/client"
	"github.com/spectre-project/spectred/cmd/spectrewallet/daemon/pb"
	"github.com/spectre-project/spectred/cmd/spectrewallet/utils"
)

func balance(conf *balanceConfig) error {
	daemonClient, tearDown, err := client.Connect(conf.DaemonAddress)
	if err != nil {
		return err
	}
	defer tearDown()

	ctx, cancel := context.WithTimeout(context.Background(), daemonTimeout)
	defer cancel()
	response, err := daemonClient.GetBalance(ctx, &pb.GetBalanceRequest{})
	if err != nil {
		return err
	}

	pendingSuffix := ""
	if response.Pending > 0 {
		pendingSuffix = " (pending)"
	}
	if conf.Verbose {
		pendingSuffix = ""
		println("Address                                                                       Available             Pending")
		println("-----------------------------------------------------------------------------------------------------------")
		for _, addressBalance := range response.AddressBalances {
			fmt.Printf("%s %s %s\n", addressBalance.Address, utils.FormatSpr(addressBalance.Available), utils.FormatSpr(addressBalance.Pending))
		}
		println("-----------------------------------------------------------------------------------------------------------")
		print("                                                 ")
	}
	fmt.Printf("Total balance, SPR %s %s%s\n", utils.FormatSpr(response.Available), utils.FormatSpr(response.Pending), pendingSuffix)

	return nil
}
