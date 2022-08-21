package keeper

import (
	"context"
	"time"

	"parrots/x/parrots/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateComment(goCtx context.Context, msg *types.MsgCreateComment) (*types.MsgCreateCommentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var comment = types.Comment{
		Creator:   msg.Creator,
		Username:  msg.Username,
		Comment:   msg.Comment,
		CreatedAt: time.Now().Unix(),
		BeakId:    msg.BeakId,
	}

	k.AddComment(ctx, comment)

	return &types.MsgCreateCommentResponse{}, nil
}
