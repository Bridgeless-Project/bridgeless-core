package keeper

import (
	bridgetypes "github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	swaptypes "github.com/Bridgeless-Project/bridgeless-core/v12/x/swap/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func (k Keeper) SetSwapSubmissions(ctx sdk.Context, submissions *bridgetypes.Submissions) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), swaptypes.Prefix(swaptypes.StoreSwapSubmissionPrefix))
	store.Set([]byte(submissions.Hash), k.cdc.MustMarshal(submissions))
}

func (k Keeper) GetSwapSubmissions(ctx sdk.Context, hash string) (bridgetypes.Submissions, bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), swaptypes.Prefix(swaptypes.StoreSwapSubmissionPrefix))
	bz := store.Get([]byte(hash))
	if bz == nil {
		return bridgetypes.Submissions{}, false
	}

	var submissions bridgetypes.Submissions
	k.cdc.MustUnmarshal(bz, &submissions)
	return submissions, true
}

func (k Keeper) SwapRequestHash(msg *swaptypes.MsgSubmitSwapTx) common.Hash {
	payload := &swaptypes.MsgSubmitSwapTx{
		Tx:         msg.Tx,
		IsBridgeTx: msg.IsBridgeTx,
	}

	return crypto.Keccak256Hash(k.cdc.MustMarshal(payload))
}

func hasSubmitter(submitters []string, submitter string) bool {
	for _, existing := range submitters {
		if existing == submitter {
			return true
		}
	}

	return false
}
