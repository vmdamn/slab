package slabone

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"slabone/testutil/sample"
	slabonesimulation "slabone/x/slabone/simulation"
	"slabone/x/slabone/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = slabonesimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateSlab = "op_weight_msg_create_slab"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateSlab int = 100

	opWeightMsgInspectSlab = "op_weight_msg_inspect_slab"
	// TODO: Determine the simulation weight value
	defaultWeightMsgInspectSlab int = 100

	opWeightMsgRevokeSlab = "op_weight_msg_revoke_slab"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRevokeSlab int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	slaboneGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&slaboneGenesis)
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

	var weightMsgCreateSlab int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateSlab, &weightMsgCreateSlab, nil,
		func(_ *rand.Rand) {
			weightMsgCreateSlab = defaultWeightMsgCreateSlab
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateSlab,
		slabonesimulation.SimulateMsgCreateSlab(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgInspectSlab int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgInspectSlab, &weightMsgInspectSlab, nil,
		func(_ *rand.Rand) {
			weightMsgInspectSlab = defaultWeightMsgInspectSlab
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgInspectSlab,
		slabonesimulation.SimulateMsgInspectSlab(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRevokeSlab int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRevokeSlab, &weightMsgRevokeSlab, nil,
		func(_ *rand.Rand) {
			weightMsgRevokeSlab = defaultWeightMsgRevokeSlab
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRevokeSlab,
		slabonesimulation.SimulateMsgRevokeSlab(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
