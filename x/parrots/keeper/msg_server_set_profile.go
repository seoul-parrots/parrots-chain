package keeper

import (
	"context"

	"parrots/x/parrots/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SetProfile(goCtx context.Context, msg *types.MsgSetProfile) (*types.MsgSetProfileResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var profile = types.Profile{
		Creator:        msg.Creator,
		Username:       msg.Username,
		DisplayName:    msg.DisplayName,
		Description:    msg.Description,
		RespectedBeaks: make([]uint64, 0),
	}

	id := k.AddProfile(ctx, profile)

	return &types.MsgSetProfileResponse{Id: id}, nil
}
