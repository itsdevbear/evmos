package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgBridgeIn{}, "synapse/BridgeIn", nil)
	cdc.RegisterConcrete(&MsgBridgeOut{}, "synapse/BridgeOut", nil)
	cdc.RegisterConcrete(&MsgSetPendingKmsAddress{}, "synapse/SetPendingKmsAddress", nil)
	cdc.RegisterConcrete(&MsgSetKmsAddress{}, "synapse/SetKmsAddress", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBridgeIn{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBridgeOut{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetPendingKmsAddress{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetKmsAddress{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
