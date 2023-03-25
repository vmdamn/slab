package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"slabone/x/slabone/types"
)

func (k msgServer) CreateSlab(goCtx context.Context, msg *types.MsgCreateSlab) (*types.MsgCreateSlabResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateSlabResponse{}, nil
}
