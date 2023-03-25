package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"slabone/x/slabone/types"
)

func (k msgServer) InspectSlab(goCtx context.Context, msg *types.MsgInspectSlab) (*types.MsgInspectSlabResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgInspectSlabResponse{}, nil
}
