package keeper

import (
	"encoding/binary"
	"errors"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lithammer/fuzzysearch/fuzzy"

	"parrots/x/parrots/types"

	"golang.org/x/exp/slices"
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

// UpdateBeak adds beak object to the blockchain.
func (k Keeper) UpdateBeak(ctx sdk.Context, beak types.Beak) error {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.BeakKey))

	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, beak.Id)

	appendedValue := k.cdc.MustMarshal(&beak)
	store.Set(byteKey, appendedValue)

	return nil
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

// GetBeaksByNameSubstring fetches beaks object from blockchain.
func (k Keeper) GetEveryBeakByNameSubstring(ctx sdk.Context, name string) ([]*types.Beak, error) {
	beaks, err := k.GetEveryBeak(ctx)
	if err != nil {
		return nil, errors.New("internal error")
	}

	var result []*types.Beak
	for _, beak := range beaks {
		if fuzzy.Match(name, beak.Name) {
			result = append(result, beak)
		}
	}

	return result, nil
}

// GetBeaksByTag fetches beaks object from blockchain.
func (k Keeper) GetEveryBeakByTag(ctx sdk.Context, tag string) ([]*types.Beak, error) {
	beaks, err := k.GetEveryBeak(ctx)
	if err != nil {
		return nil, errors.New("internal error")
	}

	var result []*types.Beak
	for _, beak := range beaks {
		idx := slices.IndexFunc(beak.Tags, func(beakTag string) bool { return beakTag == tag })
		if idx >= 0 {
			result = append(result, beak)
		}
	}

	return result, nil
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
