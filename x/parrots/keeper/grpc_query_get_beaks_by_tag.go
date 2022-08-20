package keeper

import (
	"context"

	"parrots/x/parrots/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetBeaksByTag(goCtx context.Context, req *types.QueryGetBeaksByTagRequest) (*types.QueryGetBeaksByTagResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	beaks, err := k.GetEveryBeakByTag(ctx, req.Tag)

	if err != nil {
		return nil, status.Error(codes.Internal, "Internal service error")
	}

	return &types.QueryGetBeaksByTagResponse{Beaks: beaks}, nil
}
