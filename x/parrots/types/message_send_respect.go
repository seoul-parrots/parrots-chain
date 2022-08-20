package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSendRespect = "send_respect"

var _ sdk.Msg = &MsgSendRespect{}

func NewMsgSendRespect(creator string, beakId uint64) *MsgSendRespect {
	return &MsgSendRespect{
		Creator: creator,
		BeakId:  beakId,
	}
}

func (msg *MsgSendRespect) Route() string {
	return RouterKey
}

func (msg *MsgSendRespect) Type() string {
	return TypeMsgSendRespect
}

func (msg *MsgSendRespect) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSendRespect) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendRespect) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
