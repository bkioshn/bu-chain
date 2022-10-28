package keeper_test

import (
	"strconv"
	"testing"

	keepertest "bu-chain/testutil/keeper"
	"bu-chain/testutil/nullify"
	"bu-chain/x/exchange/keeper"
	"bu-chain/x/exchange/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNExchangeRate(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ExchangeRate {
	items := make([]types.ExchangeRate, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetExchangeRate(ctx, items[i])
	}
	return items
}

func createMockExchangeRate(keeper *keeper.Keeper, ctx sdk.Context) {
	mockExchanges := []types.ExchangeRate{
		{
			Index:   "bubu-ngum",
			Rate:    "10",
			Creator: "alice",
		},
		{
			Index:   "bubu-day",
			Rate:    "0.1",
			Creator: "alice",
		},
		{
			Index:   "day-ngum",
			Rate:    "100",
			Creator: "alice",
		},
	}
	for i := range mockExchanges {
		keeper.SetExchangeRate(ctx, mockExchanges[i])
	}
}
func TestExchangeRateGet(t *testing.T) {
	keeper, ctx := keepertest.ExchangeKeeper(t)
	items := createNExchangeRate(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetExchangeRate(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestExchangeRateRemove(t *testing.T) {
	keeper, ctx := keepertest.ExchangeKeeper(t)
	items := createNExchangeRate(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveExchangeRate(ctx,
			item.Index,
		)
		_, found := keeper.GetExchangeRate(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestExchangeRateGetAll(t *testing.T) {
	keeper, ctx := keepertest.ExchangeKeeper(t)
	items := createNExchangeRate(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllExchangeRate(ctx)),
	)
}
