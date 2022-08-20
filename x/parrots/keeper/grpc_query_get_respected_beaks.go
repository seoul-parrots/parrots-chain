package keeper

import (
	"context"

	"parrots/x/parrots/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetRespectedBeaks(goCtx context.Context, req *types.QueryGetRespectedBeaksRequest) (*types.QueryGetRespectedBeaksResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	beaks, err := k.GetEveryRespectedBeaks(ctx, req.Id)

	if err != nil {
		return nil, status.Error(codes.Internal, "Internal service error")
	}

	return &types.QueryGetRespectedBeaksResponse{Beaks: beaks}, nil
}
