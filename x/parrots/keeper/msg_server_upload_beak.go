package keeper

import (
	"context"

	"parrots/x/parrots/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UploadBeak(goCtx context.Context, msg *types.MsgUploadBeak) (*types.MsgUploadBeakResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var beak = types.Beak{
		Creator:         msg.Creator,
		FileIndex:       msg.FileIndex,
		Name:            msg.Name,
		CreatorUsername: msg.CreatorUsername,
		Description:     msg.Description,
		License:         msg.License,
		Tags:            msg.Tags,
		LinkedBeaks:     msg.LinkedBeaks,
	}

	id := k.AddBeak(ctx, beak)

	return &types.MsgUploadBeakResponse{Id: id}, nil
}
