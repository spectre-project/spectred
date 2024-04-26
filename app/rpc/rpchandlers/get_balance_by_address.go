package rpchandlers

import (
	"github.com/pkg/errors"
	"github.com/spectre-project/spectred/app/appmessage"
	"github.com/spectre-project/spectred/app/rpc/rpccontext"
	"github.com/spectre-project/spectred/domain/consensus/utils/txscript"
	"github.com/spectre-project/spectred/infrastructure/network/netadapter/router"
	"github.com/spectre-project/spectred/util"
)

// HandleGetBalanceByAddress handles the respectively named RPC command
func HandleGetBalanceByAddress(context *rpccontext.Context, _ *router.Router, request appmessage.Message) (appmessage.Message, error) {
	if !context.Config.UTXOIndex {
		errorMessage := &appmessage.GetUTXOsByAddressesResponseMessage{}
		errorMessage.Error = appmessage.RPCErrorf("Method unavailable when spectred is run without --utxoindex")
		return errorMessage, nil
	}

	getBalanceByAddressRequest := request.(*appmessage.GetBalanceByAddressRequestMessage)

	balance, err := getBalanceByAddress(context, getBalanceByAddressRequest.Address)
	if err != nil {
		rpcError := &appmessage.RPCError{}
		if !errors.As(err, &rpcError) {
			return nil, err
		}
		errorMessage := &appmessage.GetUTXOsByAddressesResponseMessage{}
		errorMessage.Error = rpcError
		return errorMessage, nil
	}

	response := appmessage.NewGetBalanceByAddressResponse(balance)
	return response, nil
}

func getBalanceByAddress(context *rpccontext.Context, addressString string) (uint64, error) {
	address, err := util.DecodeAddress(addressString, context.Config.ActiveNetParams.Prefix)
	if err != nil {
		return 0, appmessage.RPCErrorf("Couldn't decode address '%s': %s", addressString, err)
	}

	scriptPublicKey, err := txscript.PayToAddrScript(address)
	if err != nil {
		return 0, appmessage.RPCErrorf("Could not create a scriptPublicKey for address '%s': %s", addressString, err)
	}
	utxoOutpointEntryPairs, err := context.UTXOIndex.UTXOs(scriptPublicKey)
	if err != nil {
		return 0, err
	}

	balance := uint64(0)
	for _, utxoOutpointEntryPair := range utxoOutpointEntryPairs {
		balance += utxoOutpointEntryPair.Amount()
	}
	return balance, nil
}
