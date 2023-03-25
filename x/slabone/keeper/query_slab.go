package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"slabone/x/slabone/types"
)

func (k Keeper) SlabAll(goCtx context.Context, req *types.QueryAllSlabRequest) (*types.QueryAllSlabResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var slabs []types.Slab
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	slabStore := prefix.NewStore(store, types.KeyPrefix(types.SlabKey))

	pageRes, err := query.Paginate(slabStore, req.Pagination, func(key []byte, value []byte) error {
		var slab types.Slab
		if err := k.cdc.Unmarshal(value, &slab); err != nil {
			return err
		}

		slabs = append(slabs, slab)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllSlabResponse{Slab: slabs, Pagination: pageRes}, nil
}

func (k Keeper) Slab(goCtx context.Context, req *types.QueryGetSlabRequest) (*types.QueryGetSlabResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	slab, found := k.GetSlab(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetSlabResponse{Slab: slab}, nil
}
