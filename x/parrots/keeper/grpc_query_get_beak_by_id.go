package keeper

import (
	"context"

	"parrots/x/parrots/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetBeakById(goCtx context.Context, req *types.QueryGetBeakByIdRequest) (*types.QueryGetBeakByIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	beak, err := k.GetSingleBeak(ctx, req.Id)

	if err != nil {
		return nil, status.Error(codes.Internal, "Internal service error")
	}

	return &types.QueryGetBeakByIdResponse{Beak: beak}, nil
}
