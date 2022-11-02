package goldoracle_test

import (
	"testing"

	keepertest "bu-chain/testutil/keeper"
	"bu-chain/testutil/nullify"
	"bu-chain/x/goldoracle"
	"bu-chain/x/goldoracle/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.GoldoracleKeeper(t)
	goldoracle.InitGenesis(ctx, *k, genesisState)
	got := goldoracle.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	// this line is used by starport scaffolding # genesis/test/assert
}
