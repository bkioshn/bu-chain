package goldoracle

import (
	"bu-chain/x/goldoracle/keeper"
	"bu-chain/x/goldoracle/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func reqGoldPrice(ctx sdk.Context, srv types.MsgServer) {
	srcPort := "goldoracle"
	srcChannel := "channel-1"
	curTime := ctx.BlockTime().UnixNano()

	msgReqGold := types.MsgSendReqGoldPriceRequest{
		Creator:   ctx.ChainID(),
		Port:      srcPort,
		ChannelID: srcChannel,
		Timeout:   uint64(curTime + 1.8e+11),
	}
	srv.SendReqGoldPrice(ctx, &msgReqGold)
}

func handleEndBlock(ctx sdk.Context, k keeper.Keeper) {
	srv := keeper.NewMsgServerImpl(k)

	reqGoldPrice(ctx, srv)
}
