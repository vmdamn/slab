package keeper

import (
	"context"
	"strconv"
	"slabone/x/slabone/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) InspectSlab(goCtx context.Context, msg *types.MsgInspectSlab) (*types.MsgInspectSlabResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	bHeight := uint64(ctx.BlockHeight())
	convbHeight := strconv.FormatUint(bHeight,10)

	bTime := (ctx.BlockTime())
    convbTime := bTime.String()

	
	slab, found := k.GetSlab(ctx, msg.Id)
	
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "key %d doesn't exist", msg.Id)		
	}

	if slab.State != "Created" {
		return nil, sdkerrors.Wrapf(types.ErrWrongSlabState, "%v", slab.State)
	}

	if slab.OriginatorSocialId == msg.VetterSocialId {
		return nil, sdkerrors.Wrapf(types.ErrVetterIsOriginator, "%s", slab.OriginatorSocialId)
	}

	if slab.OriginatorChainAddr == msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrVetterIsOriginator, "%s", slab.OriginatorChainAddr)
	}

	if slab.DirectedTowards != msg.VetterSocialId {
		return nil, sdkerrors.Wrapf(types.ErrNotDirectedTowards, "%s", slab.DirectedTowards)
	}


	slab.VetterSocialId = msg.VetterSocialId
	slab.VetterChainAddr = msg.Creator
	slab.State = "Vetted"
	slab.VettingCtxHeight = convbHeight
	slab.VettingCtxTime = convbTime
	slab.VettingNote = msg.VettingNote
	slab.UriVetter = msg.UriVetter


	k.SetSlab(ctx,slab)


	ctx.EventManager().EmitEvent(
		sdk.NewEvent(sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(types.SlabVettedEventVetter, msg.Creator),
			sdk.NewAttribute(types.SlabVettedEventVetterSocialId, msg.VetterSocialId),
			sdk.NewAttribute(types.SlabVettedEventSlabState, slab.State),
			sdk.NewAttribute(types.SlabVettedEventSlabVettingNote, slab.VettingNote),
			sdk.NewAttribute(types.SlabVettedEventSlabId, strconv.FormatUint(msg.Id,10)),
			sdk.NewAttribute(types.SlabVettedEventSlabCtxTime, slab.VettingCtxTime),
		),
	)

	return &types.MsgInspectSlabResponse{}, nil
}
