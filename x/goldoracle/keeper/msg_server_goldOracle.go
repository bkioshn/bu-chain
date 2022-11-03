package keeper

import (
	"bu-chain/x/goldoracle/types"
	"context"

	"github.com/bandprotocol/bandchain-packet/obi"
	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v5/modules/core/02-client/types"
)

type Calldata struct {
	symbol     []string
	multiplier uint64
}

const (
	clientID       = "buchain"
	oracleScriptID = 38 // Band Forex & Commodities
	askCount       = 10
	minCount       = 8
	prepareGas     = 100000
	executeGas     = 750000
)

func (k msgServer) SendReqGoldPrice(goCtx context.Context, msg *types.MsgSendReqGoldPriceRequest) (*types.MsgSendReqGoldPriceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Constructing packet
	var packet types.OracleRequestPacketData

	calldata := Calldata{
		symbol:     []string{"XAU"},
		multiplier: 1e6,
	}
	packet.ClientID = clientID
	packet.OracleScriptID = oracleScriptID
	packet.Calldata = string(obi.MustEncode(calldata))
	packet.AskCount = askCount
	packet.MinCount = minCount
	packet.FeeLimit = sdk.NewInt64Coin("uband", 100000)
	packet.PrepareGas = prepareGas
	packet.ExecuteGas = executeGas

	// Transmit the packet
	err := k.TransmitOraclePacket(
		ctx,
		packet,
		msg.Port,
		msg.ChannelID,
		clienttypes.ZeroHeight(),
		msg.Timeout,
	)
	if err != nil {
		return nil, err
	}

	return &types.MsgSendReqGoldPriceResponse{}, nil
}
