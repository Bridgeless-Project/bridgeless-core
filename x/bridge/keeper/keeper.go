package keeper

import (
	"fmt"

	"cosmossdk.io/errors"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace
		bank       types.BankKeeper
		erc20      types.ERC20Keeper
		hooks      types.BridgeHook
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	bankkeeper types.BankKeeper,
	erc20 types.ERC20Keeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
		bank:       bankkeeper,
		erc20:      erc20,
	}
}

func (k *Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k *Keeper) PartiesDistributeFee(ctx sdk.Context, epochId uint32, fee sdk.Coin) error {
	epoch, found := k.GetEpoch(ctx, epochId)
	if !found {
		return types.ErrEpochNotFound
	}

	if len(epoch.Parties) <= 0 {
		return types.ErrInvalidParty
	}

	tokensToSend := fee.Amount.QuoRaw(int64(len(epoch.Parties)))
	for _, party := range epoch.Parties {
		partyAddress, err := sdk.AccAddressFromBech32(party.Address)
		if err != nil {
			return errors.Wrap(err, "invalid party address")
		}

		err = k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, partyAddress, sdk.NewCoins(sdk.NewCoin(fee.Denom, tokensToSend)))
		if err != nil {
			return errors.Wrap(err, "failed to distribute fee to party")
		}
	}

	return nil
}

func (k *Keeper) SetHooks(hooks types.BridgeHook) *Keeper {
	if k.hooks != nil {
		panic("cannot set bridge hooks twice")
	}

	k.hooks = hooks
	return k
}
