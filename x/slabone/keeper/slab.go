package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"slabone/x/slabone/types"
)

// GetSlabCount get the total number of slab
func (k Keeper) GetSlabCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.SlabCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetSlabCount set the total number of slab
func (k Keeper) SetSlabCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.SlabCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendSlab appends a slab in the store with a new id and update the count
func (k Keeper) AppendSlab(
	ctx sdk.Context,
	slab types.Slab,
) uint64 {
	// Create the slab
	count := k.GetSlabCount(ctx)

	// Set the ID of the appended value
	slab.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SlabKey))
	appendedValue := k.cdc.MustMarshal(&slab)
	store.Set(GetSlabIDBytes(slab.Id), appendedValue)

	// Update slab count
	k.SetSlabCount(ctx, count+1)

	return count
}

// SetSlab set a specific slab in the store
func (k Keeper) SetSlab(ctx sdk.Context, slab types.Slab) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SlabKey))
	b := k.cdc.MustMarshal(&slab)
	store.Set(GetSlabIDBytes(slab.Id), b)
}

// GetSlab returns a slab from its id
func (k Keeper) GetSlab(ctx sdk.Context, id uint64) (val types.Slab, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SlabKey))
	b := store.Get(GetSlabIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveSlab removes a slab from the store
func (k Keeper) RemoveSlab(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SlabKey))
	store.Delete(GetSlabIDBytes(id))
}

// GetAllSlab returns all slab
func (k Keeper) GetAllSlab(ctx sdk.Context) (list []types.Slab) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SlabKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Slab
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetSlabIDBytes returns the byte representation of the ID
func GetSlabIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetSlabIDFromBytes returns ID in uint64 format from a byte array
func GetSlabIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
