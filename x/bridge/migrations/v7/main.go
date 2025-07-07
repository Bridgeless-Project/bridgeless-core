package v7

import (
	"fmt"
	oldTypes "github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/migrations/v7/types"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	ctx.Logger().Info(fmt.Sprintf("Performing v12.1.20-rc1 %s module migrations", types.ModuleName))

	chainsStore := prefix.NewStore(ctx.KVStore(storeKey), types.Prefix(types.StoreChainPrefix))
	iterator := sdk.KVStorePrefixIterator(chainsStore, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var oldChain oldTypes.Chain
		cdc.MustUnmarshal(iterator.Value(), &oldChain)

		newChain := types.Chain{
			Id:            oldChain.Id,
			Type:          types.ChainType(oldChain.Type),
			BridgeAddress: oldChain.BridgeAddress,
			Operator:      oldChain.Operator,
			Confirmations: 0,
			Name:          "",
		}

		// set new chain instead of old one
		chainsStore.Set(
			[]byte(oldChain.Id),
			cdc.MustMarshal(&newChain),
		)
	}

	return nil
}
