package keeper

import (
	"context"
	"strconv"
	"slabone/x/slabone/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) RevokeSlab(goCtx context.Context, msg *types.MsgRevokeSlab) (*types.MsgRevokeSlabResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	bHeight := uint64(ctx.BlockHeight())
	convbHeight := strconv.FormatUint(bHeight,10)

	bTime := (ctx.BlockTime())
    convbTime := bTime.String()


	slab, found := k.GetSlab(ctx, msg.Id)
	
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "key %d doesn't exist", msg.Id)		
	}

	if slab.State != "Vetted" {
		return nil, sdkerrors.Wrapf(types.ErrWrongSlabState, "%v", slab.State)
	}

	if slab.VetterSocialId != msg.RevokerSocialId {
		return nil, sdkerrors.Wrapf(types.ErrRevokerShouldBeVetter, "%s", slab.VetterSocialId)
	}

	if slab.VetterChainAddr != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrRevokerShouldBeVetter, "%s", slab.VetterChainAddr)
	}

	

	slab.RevokerSocialId = msg.RevokerSocialId
	slab.RevokerChainAddr = msg.Creator
	slab.State = "Revoked"
	slab.RevokingCtxHeight = convbHeight
	slab.RevokingCtxTime = convbTime
	slab.RevokingNote = msg.RevokingNote
	slab.UriRevoker = msg.UriRevoker


	k.SetSlab(ctx,slab)


	ctx.EventManager().EmitEvent(
		sdk.NewEvent(sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(types.SlabRevokedEventRevoker, msg.Creator),
			sdk.NewAttribute(types.SlabRevokedEventRevokerSocialId, msg.RevokerSocialId),
			sdk.NewAttribute(types.SlabRevokedEventSlabState, slab.State),
			sdk.NewAttribute(types.SlabRevokedEventSlabRevokingNote, slab.RevokingNote),
			sdk.NewAttribute(types.SlabRevokedEventSlabId, strconv.FormatUint(msg.Id,10)),
			sdk.NewAttribute(types.SlabRevokedEventSlabCtxTime, slab.RevokingCtxTime),
		),
	)

	return &types.MsgRevokeSlabResponse{}, nil
}
