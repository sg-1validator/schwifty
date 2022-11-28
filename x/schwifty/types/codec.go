package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateCollection{}, "schwifty/CreateCollection", nil)
	cdc.RegisterConcrete(&MsgUpdateCollection{}, "schwifty/UpdateCollection", nil)
	cdc.RegisterConcrete(&MsgDeleteCollection{}, "schwifty/DeleteCollection", nil)
	cdc.RegisterConcrete(&MsgCreateNft{}, "schwifty/CreateNft", nil)
	cdc.RegisterConcrete(&MsgUpdateNft{}, "schwifty/UpdateNft", nil)
	cdc.RegisterConcrete(&MsgDeleteNft{}, "schwifty/DeleteNft", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateCollection{},
		&MsgUpdateCollection{},
		&MsgDeleteCollection{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateNft{},
		&MsgUpdateNft{},
		&MsgDeleteNft{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
