package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"parrots/x/parrots/types"
)

// AddBeak adds beak object to the blockchain.
func (k Keeper) AddComment(ctx sdk.Context, comment types.Comment) uint64 {
	count := k.GetCommentCount(ctx)

	comment.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.CommentKey))

	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, comment.Id)

	appendedValue := k.cdc.MustMarshal(&comment)
	store.Set(byteKey, appendedValue)

	k.SetCommentCount(ctx, count+1)

	return count
}

// SetBeakCount sets beak count from blockchain.
func (k Keeper) SetCommentCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.CommentCountKey))

	byteKey := []byte(types.CommentCountKey)

	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)

	store.Set(byteKey, bz)
}

// GetBeakCount retrieves beak count from blockchain.
func (k Keeper) GetCommentCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.CommentCountKey))

	byteKey := []byte(types.CommentCountKey)
	bz := store.Get(byteKey)

	if bz == nil {
		return 0
	}

	return binary.BigEndian.Uint64(bz)
}
