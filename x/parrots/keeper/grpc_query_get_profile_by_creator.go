package keeper

import (
	"context"

	"parrots/x/parrots/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetProfileByCreator(goCtx context.Context, req *types.QueryGetProfileByCreatorRequest) (*types.QueryGetProfileByCreatorResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	profile, err := k.GetSingleProfileByCreator(ctx, req.Creator)
	if err != nil {
		return nil, err
	}

	return &types.QueryGetProfileByCreatorResponse{Profile: profile}, nil
}
