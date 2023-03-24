package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "slabone/testutil/keeper"
	"slabone/x/slabone/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.SlaboneKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
