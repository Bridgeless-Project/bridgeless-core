package types

import (
	v19 "github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/versions/v19"
	v21 "github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/versions/v21"
	v24 "github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/versions/v24"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	// this line is used by starport scaffolding # 1
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRemoveTokenInfo{},
		&MsgAddTokenInfo{},

		&MsgSubmitTransactions{},
		&MsgRemoveTransaction{},
		&MsgUpdateTransaction{},

		&MsgInsertChain{},
		&MsgDeleteChain{},

		&MsgDeleteToken{},
		&MsgInsertToken{},

		&v19.MsgAddTokenInfo{},
		&v19.MsgInsertChain{},
		&v19.MsgInsertToken{},
		&v19.MsgRemoveTokenInfo{},
		&v19.MsgSubmitTransactions{},
		&v19.MsgUpdateToken{},

		&v21.MsgAddTokenInfo{},
		&v21.MsgInsertToken{},

		&v24.MsgAddTokenInfo{},
		&v24.MsgInsertToken{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
