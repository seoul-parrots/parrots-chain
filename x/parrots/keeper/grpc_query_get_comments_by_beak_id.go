package keeper

import (
	"context"

	"parrots/x/parrots/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetCommentsByBeakId(goCtx context.Context, req *types.QueryGetCommentsByBeakIdRequest) (*types.QueryGetCommentsByBeakIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	comments, err := k.GetEveryCommentByBeakId(ctx, req.BeakId)
	if err != nil {
		return nil, err
	}

	return &types.QueryGetCommentsByBeakIdResponse{Comments: comments}, nil
}
