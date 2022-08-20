package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"parrots/x/parrots/types"
)

// AddProfile adds profile object to the blockchain.
func (k Keeper) AddProfile(ctx sdk.Context, profile types.Profile) uint64 {
	count := k.GetProfileCount(ctx)

	profile.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.ProfileKey))

	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, profile.Id)

	appendedValue := k.cdc.MustMarshal(&profile)
	store.Set(byteKey, appendedValue)

	k.SetProfileCount(ctx, count+1)

	return count
}

// GetSingleProfile fetches profile object from blockchain.
func (k Keeper) GetSingleProfile(ctx sdk.Context, id uint64) (*types.Profile, error) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.ProfileKey))

	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, id)

	profile := &types.Profile{}
	bz := store.Get(byteKey)
	if err := k.cdc.Unmarshal(bz, profile); err != nil {
		return nil, err
	}

	return profile, nil
}

// SetProfileCount sets profile count from blockchain.
func (k Keeper) SetProfileCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.ProfileCountKey))

	byteKey := []byte(types.ProfileCountKey)

	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)

	store.Set(byteKey, bz)
}

// GetProfileCount retrieves profile count from blockchain.
func (k Keeper) GetProfileCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.ProfileCountKey))

	byteKey := []byte(types.ProfileCountKey)
	bz := store.Get(byteKey)

	if bz == nil {
		return 0
	}

	return binary.BigEndian.Uint64(bz)
}
