package keeper

import (
	"bu-chain/x/goldoracle/types"
	"context"

	"github.com/bandprotocol/chain/v2/pkg/obi"

	pkg "github.com/bandprotocol/bandchain-packet/packet"
	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v5/modules/core/02-client/types"
)

type Calldata struct {
	Symbol     []string
	Multiplier uint64
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
	calldata := Calldata{
		Symbol:     []string{"XAU"},
		Multiplier: 1e6,
	}

	coins := sdk.NewCoins(sdk.NewCoin("uband", sdk.NewInt(100000)))
	packet := pkg.NewOracleRequestPacketData(clientID, oracleScriptID, obi.MustEncode(calldata), askCount, minCount, coins, prepareGas, executeGas)

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
