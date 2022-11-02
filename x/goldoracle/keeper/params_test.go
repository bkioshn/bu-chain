package keeper_test

import (
	"testing"

	testkeeper "bu-chain/testutil/keeper"
	"bu-chain/x/goldoracle/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.GoldoracleKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
