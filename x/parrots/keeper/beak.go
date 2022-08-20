package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"parrots/x/parrots/types"
)

// AddBeak adds beak object to the blockchain.
func (k Keeper) AddBeak(ctx sdk.Context, beak types.Beak) uint64 {
	count := k.GetBeakCount(ctx)

	beak.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.BeakKey))

	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, beak.Id)

	appendedValue := k.cdc.MustMarshal(&beak)
	store.Set(byteKey, appendedValue)

	k.SetBeakCount(ctx, count+1)

	return count
}

// GetEveryBeak fetches every beak object from blockchain.
func (k Keeper) GetEveryBeak(ctx sdk.Context) ([]*types.Beak, error) {
	var beaks []*types.Beak
	count := k.GetBeakCount(ctx)
	for i := uint64(0); i < count; i++ {
		beak, err := k.GetSingleBeak(ctx, i)
		if err != nil {
			return nil, err
		}

		beaks = append(beaks, beak)
	}

	return beaks, nil
}

// GetSingleBeak fetches beak object from blockchain.
func (k Keeper) GetSingleBeak(ctx sdk.Context, id uint64) (*types.Beak, error) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.BeakKey))

	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, id)

	beak := &types.Beak{}
	bz := store.Get(byteKey)
	if err := k.cdc.Unmarshal(bz, beak); err != nil {
		return nil, err
	}

	return beak, nil
}

// SetBeakCount sets beak count from blockchain.
func (k Keeper) SetBeakCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.BeakCountKey))

	byteKey := []byte(types.BeakCountKey)

	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)

	store.Set(byteKey, bz)
}

// GetBeakCount retrieves beak count from blockchain.
func (k Keeper) GetBeakCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.BeakCountKey))

	byteKey := []byte(types.BeakCountKey)
	bz := store.Get(byteKey)

	if bz == nil {
		return 0
	}

	return binary.BigEndian.Uint64(bz)
}
