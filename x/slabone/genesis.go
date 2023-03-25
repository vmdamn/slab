package slabone

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"slabone/x/slabone/keeper"
	"slabone/x/slabone/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the slab
	for _, elem := range genState.SlabList {
		k.SetSlab(ctx, elem)
	}

	// Set slab count
	k.SetSlabCount(ctx, genState.SlabCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.SlabList = k.GetAllSlab(ctx)
	genesis.SlabCount = k.GetSlabCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
