package keeper

import (
	"context"
	"errors"

	"parrots/x/parrots/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SendRespect(goCtx context.Context, msg *types.MsgSendRespect) (*types.MsgSendRespectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.AddRespect(ctx, msg.Creator, msg.BeakId); err != nil {
		return nil, errors.New("Failed to show respect")
	}

	return &types.MsgSendRespectResponse{}, nil
}
