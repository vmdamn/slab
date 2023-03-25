package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"slabone/x/slabone/types"
	"strconv"
)

func (k msgServer) CreateSlab(goCtx context.Context, msg *types.MsgCreateSlab) (*types.MsgCreateSlabResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)


	bHeight := uint64(ctx.BlockHeight())
	convbHeight := strconv.FormatUint(bHeight,10)

	bTime := (ctx.BlockTime())
    convbTime := bTime.String()

	var slab = types.Slab{
		OriginatorSocialId: msg.OriginatorSocialId,
		OriginatorChainAddr:msg.Creator,
		DirectedTowards: 	msg.DirectedTowards,
		Assertion:		   	msg.Assertion,
		UriOriginator:		msg.UriOriginator,
		State:				"Created",
		OriginatedCtxHeight:convbHeight,
		OriginatedCtxTime: 	convbTime,
		VetterSocialId:		"",
		VetterChainAddr:	"",
		VettingCtxHeight:	"",
		VettingCtxTime: 	"",
		VettingNote: 		"",
		UriVetter:			"",
		RevokingCtxHeight:	"",
		RevokingCtxTime:	"",
		RevokingNote:		"",
		UriRevoker:			"",
	}
	
	id := k.AppendSlab(ctx, slab)
	

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(types.SlabCreatedEventOriginator, msg.Creator),
			sdk.NewAttribute(types.SlabCreatedEventOriginatorSocialId, msg.OriginatorSocialId),
			sdk.NewAttribute(types.SlabCreatedEventAssertion, msg.Assertion),
			sdk.NewAttribute(types.SlabCreatedEventDirectedTowards, msg.DirectedTowards),
			sdk.NewAttribute(types.SlabCreatedEventSlabState, slab.State),
			sdk.NewAttribute(types.SlabCreatedEventSlabId, strconv.FormatUint(id,10)),
			sdk.NewAttribute(types.SlabCreatedEventSlabCtxTime, slab.OriginatedCtxTime),
		),
	)



	return &types.MsgCreateSlabResponse{Id: id}, nil
}
