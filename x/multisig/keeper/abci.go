package keeper

import (
	"fmt"

	"github.com/Bridgeless-Project/bridgeless-core/v12/x/multisig/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) EndBlocker(ctx sdk.Context) {
	currentBlock := uint64(ctx.BlockHeight())

	k.IterateProposals(ctx, func(proposal types.Proposal) (stop bool) {
		if proposal.VotingEndBlock == currentBlock {
			proposal.FinalTallyResult = k.Tally(ctx, proposal.Id)
			proposal.Status = types.ProposalStatus_REJECTED

			group, _ := k.GetGroup(ctx, proposal.Group)

			if proposal.FinalTallyResult.YesCount >= group.Threshold {
				proposal.Status = types.ProposalStatus_ACCEPTED
			}

			k.SetProposal(ctx, proposal)

			ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeProposalStatusChanged,
				sdk.NewAttribute(types.AttributeKeyProposal, fmt.Sprintf("%d", proposal.Id)),
				sdk.NewAttribute(types.AttributeKeyProposalStatus, proposal.Status.String()),
			))
		}

		if proposal.Status == types.ProposalStatus_ACCEPTED {
			var err error
			logs := fmt.Sprintf("proposal execution successful on proposal %d", proposal.Id)
			proposal.Status, err = k.ExecuteProposal(ctx, proposal)
			if err != nil {
				logs = fmt.Sprintf("proposal execution failed on proposal %d, because of error %s", proposal.Id, err.Error())
			}

			k.SetProposal(ctx, proposal)

			ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeProposalExecuted,
				sdk.NewAttribute(types.AttributeKeyProposal, fmt.Sprintf("%d", proposal.Id)),
				sdk.NewAttribute(types.AttributeKeyProposalStatus, proposal.Status.String()),
				sdk.NewAttribute(types.AttributeKeyProposalExecutionLogs, logs),
			))
		}

		return false
	})
}
