package exchange_test

import (
	"testing"

	keepertest "bu-chain/testutil/keeper"
	"bu-chain/testutil/nullify"
	"bu-chain/x/exchange"
	"bu-chain/x/exchange/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		ExchangeRateList: []types.ExchangeRate{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ExchangeKeeper(t)
	exchange.InitGenesis(ctx, *k, genesisState)
	got := exchange.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.ExchangeRateList, got.ExchangeRateList)
	// this line is used by starport scaffolding # genesis/test/assert
}
