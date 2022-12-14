package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgSetProfile{}, "parrots/SetProfile", nil)
	cdc.RegisterConcrete(&MsgUploadBeak{}, "parrots/UploadBeak", nil)
	cdc.RegisterConcrete(&MsgSendRespect{}, "parrots/SendRespect", nil)
	cdc.RegisterConcrete(&MsgCreateComment{}, "parrots/CreateComment", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetProfile{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUploadBeak{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendRespect{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateComment{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
