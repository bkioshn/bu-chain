package exchange

import (
	"math/rand"

	"bu-chain/testutil/sample"
	exchangesimulation "bu-chain/x/exchange/simulation"
	"bu-chain/x/exchange/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = exchangesimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgExchangeToken = "op_weight_msg_exchange_token"
	// TODO: Determine the simulation weight value
	defaultWeightMsgExchangeToken int = 100

	opWeightMsgCreateExchangeRate = "op_weight_msg_exchange_rate"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateExchangeRate int = 100

	opWeightMsgUpdateExchangeRate = "op_weight_msg_exchange_rate"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateExchangeRate int = 100

	opWeightMsgDeleteExchangeRate = "op_weight_msg_exchange_rate"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteExchangeRate int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	exchangeGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		ExchangeRateList: []types.ExchangeRate{
			{
				Creator: sample.AccAddress(),
				Index:   "0",
			},
			{
				Creator: sample.AccAddress(),
				Index:   "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&exchangeGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgExchangeToken int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgExchangeToken, &weightMsgExchangeToken, nil,
		func(_ *rand.Rand) {
			weightMsgExchangeToken = defaultWeightMsgExchangeToken
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgExchangeToken,
		exchangesimulation.SimulateMsgExchangeToken(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateExchangeRate int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateExchangeRate, &weightMsgCreateExchangeRate, nil,
		func(_ *rand.Rand) {
			weightMsgCreateExchangeRate = defaultWeightMsgCreateExchangeRate
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateExchangeRate,
		exchangesimulation.SimulateMsgCreateExchangeRate(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateExchangeRate int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateExchangeRate, &weightMsgUpdateExchangeRate, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateExchangeRate = defaultWeightMsgUpdateExchangeRate
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateExchangeRate,
		exchangesimulation.SimulateMsgUpdateExchangeRate(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteExchangeRate int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteExchangeRate, &weightMsgDeleteExchangeRate, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteExchangeRate = defaultWeightMsgDeleteExchangeRate
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteExchangeRate,
		exchangesimulation.SimulateMsgDeleteExchangeRate(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
