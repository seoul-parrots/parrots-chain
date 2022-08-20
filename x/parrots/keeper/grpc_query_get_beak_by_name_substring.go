package keeper

import (
	"context"

	"parrots/x/parrots/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetBeaksByNameSubstring(goCtx context.Context, req *types.QueryGetBeaksByNameSubstringRequest) (*types.QueryGetBeaksByNameSubstringResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	beaks, err := k.GetEveryBeakByNameSubstring(ctx, req.NameSubstring)

	if err != nil {
		return nil, status.Error(codes.Internal, "Internal service error")
	}

	return &types.QueryGetBeaksByNameSubstringResponse{Beaks: beaks}, nil
}
