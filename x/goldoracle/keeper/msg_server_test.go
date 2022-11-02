package keeper_test

import (
	"context"
	"testing"

	keepertest "bu-chain/testutil/keeper"
	"bu-chain/x/goldoracle/keeper"
	"bu-chain/x/goldoracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.GoldoracleKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
