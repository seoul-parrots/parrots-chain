package keeper

import (
	"encoding/binary"
	"errors"

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

// Update Profile modifies profile object to the blockchain.
func (k Keeper) UpdateProfile(ctx sdk.Context, profile *types.Profile) error {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.ProfileKey))

	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, profile.Id)

	appendedValue := k.cdc.MustMarshal(profile)
	store.Set(byteKey, appendedValue)

	return nil
}

// GetEveryProfile fetches every profile object from blockchain.
func (k Keeper) GetEveryProfile(ctx sdk.Context) ([]*types.Profile, error) {
	var profiles []*types.Profile
	count := k.GetProfileCount(ctx)
	for i := uint64(0); i < count; i++ {
		profile, err := k.GetSingleProfile(ctx, i)
		if err != nil {
			return nil, err
		}

		profiles = append(profiles, profile)
	}

	return profiles, nil
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

// GetSingleProfileByUsername fetches profile object using username from blockchain.
func (k Keeper) GetSingleProfileByUsername(ctx sdk.Context, username string) (*types.Profile, error) {
	profiles, err := k.GetEveryProfile(ctx)
	if err != nil {
		return nil, errors.New("internal error")
	}

	for _, profile := range profiles {
		if profile.Username == username {
			return profile, nil
		}
	}
	return nil, errors.New("not found by given username")
}

// GetSingleProfileByCreator fetches profile object using creator from blockchain.
func (k Keeper) GetSingleProfileByCreator(ctx sdk.Context, creator string) (*types.Profile, error) {
	profiles, err := k.GetEveryProfile(ctx)
	if err != nil {
		return nil, errors.New("internal error")
	}

	for _, profile := range profiles {
		if profile.Creator == creator {
			return profile, nil
		}
	}
	return nil, errors.New("not found by given creator")
}

// GetSingleProfile fetches profile object from blockchain.
func (k Keeper) GetEveryRespectedBeaks(ctx sdk.Context, id uint64) ([]*types.Beak, error) {
	profile, err := k.GetSingleProfile(ctx, id)
	if err != nil {
		return nil, errors.New("internal error")
	}

	var beaks []*types.Beak
	for _, beakId := range profile.RespectedBeaks {
		beak, err := k.GetSingleBeak(ctx, beakId)
		if err != nil {
			return nil, errors.New("internal error")
		}
		beaks = append(beaks, beak)
	}

	return beaks, nil
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
