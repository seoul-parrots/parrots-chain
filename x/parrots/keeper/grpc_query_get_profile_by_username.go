package keeper

import (
	"context"

	"parrots/x/parrots/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetProfileByUsername(goCtx context.Context, req *types.QueryGetProfileByUsernameRequest) (*types.QueryGetProfileByUsernameResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	profile, err := k.GetSingleProfileByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}

	return &types.QueryGetProfileByUsernameResponse{Profile: profile}, nil
}
