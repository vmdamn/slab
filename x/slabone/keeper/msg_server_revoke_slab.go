package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"slabone/x/slabone/types"
)

func (k msgServer) RevokeSlab(goCtx context.Context, msg *types.MsgRevokeSlab) (*types.MsgRevokeSlabResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgRevokeSlabResponse{}, nil
}
