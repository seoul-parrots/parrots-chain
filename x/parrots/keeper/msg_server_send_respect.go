package keeper

import (
	"context"

	"parrots/x/parrots/types"
)

func (k msgServer) SendRespect(goCtx context.Context, msg *types.MsgSendRespect) (*types.MsgSendRespectResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)

	// var beakId := msg.beakId
	// respectCount, err = k.AddRespect(ctx, beakId)

	return &types.MsgSendRespectResponse{RespectCount: uint64(1)}, nil
}
