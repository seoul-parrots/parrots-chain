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

func (k Keeper) Profiles(goCtx context.Context, req *types.QueryProfilesRequest) (*types.QueryProfilesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	var profiles []*types.Profile

	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.ProfileKey))

	pageResponse, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
		var profile types.Profile
		if err := k.cdc.Unmarshal(value, &profile); err != nil {
			return err
		}

		profiles = append(profiles, &profile)

		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryProfilesResponse{Profile: profiles, Pagination: pageResponse}, nil
}
