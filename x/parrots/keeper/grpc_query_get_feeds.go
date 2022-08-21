package keeper

import (
	"context"
	"sort"

	"parrots/x/parrots/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetFeeds(goCtx context.Context, req *types.QueryGetFeedsRequest) (*types.QueryGetFeedsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	var feeds []*types.Feed

	beaks, err := k.GetEveryBeak(ctx)
	if err != nil {
		return nil, err
	}
	comments, err := k.GetEveryComment(ctx)
	if err != nil {
		return nil, err
	}

	for _, beak := range beaks {
		feeds = append(feeds, &types.Feed{Type: "Beak", Id: beak.Id, CreatedAt: beak.CreatedAt})
	}
	for _, comment := range comments {
		feeds = append(feeds, &types.Feed{Type: "Comment", Id: comment.Id, CreatedAt: comment.CreatedAt})
	}

	sort.Slice(feeds, func(i, j int) bool {
		return feeds[i].CreatedAt > feeds[j].CreatedAt
	})

	return &types.QueryGetFeedsResponse{Feeds: feeds}, nil
}
