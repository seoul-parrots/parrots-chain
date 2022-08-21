package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUploadBeak = "upload_beak"

var _ sdk.Msg = &MsgUploadBeak{}

func NewMsgUploadBeak(creator string, fileIndex string, name string, creatorUsername string, creatorDisplayName string, description string, license string, linkedBeaks []uint64, tags []string) *MsgUploadBeak {
	return &MsgUploadBeak{
		Creator:            creator,
		FileIndex:          fileIndex,
		Name:               name,
		CreatorUsername:    creatorUsername,
		CreatorDisplayName: creatorDisplayName,
		Description:        description,
		License:            license,
		LinkedBeaks:        linkedBeaks,
		Tags:               tags,
	}
}

func (msg *MsgUploadBeak) Route() string {
	return RouterKey
}

func (msg *MsgUploadBeak) Type() string {
	return TypeMsgUploadBeak
}

func (msg *MsgUploadBeak) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUploadBeak) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUploadBeak) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
