package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"parrots/x/parrots/types"
)

// AddRespect adds respect object to the blockchain.
func (k Keeper) AddRespect(ctx sdk.Context, creator string, beakId uint64) error {
	respect_count := k.GetRespectCount(ctx, beakId)
	k.SetRespectCount(ctx, beakId, respect_count+1)

	profile, err := k.GetSingleProfileByCreator(ctx, creator)
	if err != nil {
		return err
	}

	profile.RespectedBeaks = append(profile.RespectedBeaks, beakId)
	if err := k.UpdateProfile(ctx, profile); err != nil {
		return err
	}

	return nil
}

// SetRespectCount sets profile count from blockchain.
func (k Keeper) SetRespectCount(ctx sdk.Context, beakId uint64, count uint64) error {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.BeakKey))

	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, beakId)

	beak := &types.Beak{}
	bz := store.Get(byteKey)
	if err := k.cdc.Unmarshal(bz, beak); err != nil {
		return err
	}

	beak.RespectCount += 1

	values := k.cdc.MustMarshal(beak)
	store.Set(byteKey, values)

	return nil
}

// GetRespectCount retrieves profile count from blockchain.
func (k Keeper) GetRespectCount(ctx sdk.Context, beakId uint64) uint64 {
	beak, err := k.GetSingleBeak(ctx, beakId)
	if err != nil {
		return 0
	}

	return beak.RespectCount
}
