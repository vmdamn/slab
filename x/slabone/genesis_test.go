package slabone_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "slabone/testutil/keeper"
	"slabone/testutil/nullify"
	"slabone/x/slabone"
	"slabone/x/slabone/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		SlabList: []types.Slab{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		SlabCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SlaboneKeeper(t)
	slabone.InitGenesis(ctx, *k, genesisState)
	got := slabone.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.SlabList, got.SlabList)
	require.Equal(t, genesisState.SlabCount, got.SlabCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
