package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/spectre-project/spectred/infrastructure/network/netadapter/server/grpcserver/protowire"
)

var commandTypes = []reflect.Type{
	reflect.TypeOf(protowire.SpectredMessage_AddPeerRequest{}),
	reflect.TypeOf(protowire.SpectredMessage_GetConnectedPeerInfoRequest{}),
	reflect.TypeOf(protowire.SpectredMessage_GetPeerAddressesRequest{}),
	reflect.TypeOf(protowire.SpectredMessage_GetCurrentNetworkRequest{}),
	reflect.TypeOf(protowire.SpectredMessage_GetInfoRequest{}),

	reflect.TypeOf(protowire.SpectredMessage_GetBlockRequest{}),
	reflect.TypeOf(protowire.SpectredMessage_GetBlocksRequest{}),
	reflect.TypeOf(protowire.SpectredMessage_GetHeadersRequest{}),
	reflect.TypeOf(protowire.SpectredMessage_GetBlockCountRequest{}),
	reflect.TypeOf(protowire.SpectredMessage_GetBlockDagInfoRequest{}),
	reflect.TypeOf(protowire.SpectredMessage_GetSelectedTipHashRequest{}),
	reflect.TypeOf(protowire.SpectredMessage_GetVirtualSelectedParentBlueScoreRequest{}),
	reflect.TypeOf(protowire.SpectredMessage_GetVirtualSelectedParentChainFromBlockRequest{}),
	reflect.TypeOf(protowire.SpectredMessage_ResolveFinalityConflictRequest{}),
	reflect.TypeOf(protowire.SpectredMessage_EstimateNetworkHashesPerSecondRequest{}),

	reflect.TypeOf(protowire.SpectredMessage_GetBlockTemplateRequest{}),
	reflect.TypeOf(protowire.SpectredMessage_SubmitBlockRequest{}),

	reflect.TypeOf(protowire.SpectredMessage_GetMempoolEntryRequest{}),
	reflect.TypeOf(protowire.SpectredMessage_GetMempoolEntriesRequest{}),
	reflect.TypeOf(protowire.SpectredMessage_GetMempoolEntriesByAddressesRequest{}),

	reflect.TypeOf(protowire.SpectredMessage_SubmitTransactionRequest{}),

	reflect.TypeOf(protowire.SpectredMessage_GetUtxosByAddressesRequest{}),
	reflect.TypeOf(protowire.SpectredMessage_GetBalanceByAddressRequest{}),
	reflect.TypeOf(protowire.SpectredMessage_GetCoinSupplyRequest{}),

	reflect.TypeOf(protowire.SpectredMessage_BanRequest{}),
	reflect.TypeOf(protowire.SpectredMessage_UnbanRequest{}),
}

type commandDescription struct {
	name       string
	parameters []*parameterDescription
	typeof     reflect.Type
}

type parameterDescription struct {
	name   string
	typeof reflect.Type
}

func commandDescriptions() []*commandDescription {
	commandDescriptions := make([]*commandDescription, len(commandTypes))

	for i, commandTypeWrapped := range commandTypes {
		commandType := unwrapCommandType(commandTypeWrapped)

		name := strings.TrimSuffix(commandType.Name(), "RequestMessage")
		numFields := commandType.NumField()

		var parameters []*parameterDescription
		for i := 0; i < numFields; i++ {
			field := commandType.Field(i)

			if !isFieldExported(field) {
				continue
			}

			parameters = append(parameters, &parameterDescription{
				name:   field.Name,
				typeof: field.Type,
			})
		}
		commandDescriptions[i] = &commandDescription{
			name:       name,
			parameters: parameters,
			typeof:     commandTypeWrapped,
		}
	}

	return commandDescriptions
}

func (cd *commandDescription) help() string {
	sb := &strings.Builder{}
	sb.WriteString(cd.name)
	for _, parameter := range cd.parameters {
		_, _ = fmt.Fprintf(sb, " [%s]", parameter.name)
	}
	return sb.String()
}
