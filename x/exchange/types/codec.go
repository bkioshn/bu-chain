package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgExchangeToken{}, "exchange/ExchangeToken", nil)
	cdc.RegisterConcrete(&MsgCreateExchangeRate{}, "exchange/CreateExchangeRate", nil)
	cdc.RegisterConcrete(&MsgUpdateExchangeRate{}, "exchange/UpdateExchangeRate", nil)
	cdc.RegisterConcrete(&MsgDeleteExchangeRate{}, "exchange/DeleteExchangeRate", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgExchangeToken{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateExchangeRate{},
		&MsgUpdateExchangeRate{},
		&MsgDeleteExchangeRate{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
