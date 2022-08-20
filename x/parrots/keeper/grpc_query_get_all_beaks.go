package keeper

import (
	"context"

	"parrots/x/parrots/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetAllBeaks(goCtx context.Context, req *types.QueryGetAllBeaksRequest) (*types.QueryGetAllBeaksResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	var beaks []*types.Beak

	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.BeakKey))

	pageResponse, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
		var beak types.Beak
		if err := k.cdc.Unmarshal(value, &beak); err != nil {
			return err
		}

		beaks = append(beaks, &beak)

		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryGetAllBeaksResponse{Beaks: beaks, Pagination: pageResponse}, nil
}
