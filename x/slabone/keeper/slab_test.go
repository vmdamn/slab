package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "slabone/testutil/keeper"
	"slabone/testutil/nullify"
	"slabone/x/slabone/keeper"
	"slabone/x/slabone/types"
)

func createNSlab(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Slab {
	items := make([]types.Slab, n)
	for i := range items {
		items[i].Id = keeper.AppendSlab(ctx, items[i])
	}
	return items
}

func TestSlabGet(t *testing.T) {
	keeper, ctx := keepertest.SlaboneKeeper(t)
	items := createNSlab(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetSlab(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestSlabRemove(t *testing.T) {
	keeper, ctx := keepertest.SlaboneKeeper(t)
	items := createNSlab(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveSlab(ctx, item.Id)
		_, found := keeper.GetSlab(ctx, item.Id)
		require.False(t, found)
	}
}

func TestSlabGetAll(t *testing.T) {
	keeper, ctx := keepertest.SlaboneKeeper(t)
	items := createNSlab(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllSlab(ctx)),
	)
}

func TestSlabCount(t *testing.T) {
	keeper, ctx := keepertest.SlaboneKeeper(t)
	items := createNSlab(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetSlabCount(ctx))
}
