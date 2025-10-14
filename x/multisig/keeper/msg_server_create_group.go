package keeper

import (
	"context"
	"encoding/binary"

	"github.com/Bridgeless-Project/bridgeless-core/v12/x/multisig/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

func (k msgServer) CreateGroup(goCtx context.Context, msg *types.MsgCreateGroup) (*types.MsgCreateGroupResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := k.GetParams(ctx)

	var groupAccAddr sdk.AccAddress
	// loop here in the rare case where an ADR-028-derived address creates a
	// collision with an existing address.
	i := params.GroupSequence + 1
	for {
		buf := make([]byte, 8)
		binary.BigEndian.PutUint64(buf, i)

		parentAcc := address.Module(types.ModuleName, []byte{types.GroupAccountKey})
		groupAccAddr = address.Derive(parentAcc, buf)

		account := k.accountKeeper.GetAccount(ctx, groupAccAddr)

		if account != nil {
			// handle a rare collision, in which case we just go on to the
			// next sequence value and derive a new address.
			i++
			continue
		}

		acc := k.accountKeeper.NewAccount(ctx, &authtypes.ModuleAccount{
			BaseAccount: &authtypes.BaseAccount{
				Address: groupAccAddr.String(),
			},
			Name: groupAccAddr.String(),
		})

		k.accountKeeper.SetAccount(ctx, acc)
		break
	}

	k.SetGroup(ctx, types.Group{
		Account:   groupAccAddr.String(),
		Members:   msg.Members,
		Threshold: msg.Threshold,
	})

	params.GroupSequence = i
	k.SetParams(ctx, params)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeCreateGroup,
			sdk.NewAttribute(types.AttributeKeyGroup, groupAccAddr.String()),
		),
	)

	return &types.MsgCreateGroupResponse{
		Group: groupAccAddr.String(),
	}, nil
}
