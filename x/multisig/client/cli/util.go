package cli

import (
	"encoding/json"
	"os"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type proposal struct {
	// Msgs defines an array of sdk.Msgs proto-JSON-encoded as Anys.
	Creator  string            `json:"creator"`
	Group    string            `json:"group"`
	Messages []json.RawMessage `json:"messages,omitempty"`
}

func parseSubmitProposal(cdc codec.Codec, path string) (string, string, []sdk.Msg, error) {
	var proposal proposal

	contents, err := os.ReadFile(path)
	if err != nil {
		return "", "", nil, err
	}

	err = json.Unmarshal(contents, &proposal)
	if err != nil {
		return "", "", nil, err
	}

	msgs := make([]sdk.Msg, len(proposal.Messages))
	for i, anyJSON := range proposal.Messages {
		var msg sdk.Msg
		err := cdc.UnmarshalInterfaceJSON(anyJSON, &msg)
		if err != nil {
			return "", "", nil, err
		}

		msgs[i] = msg
	}

	return proposal.Creator, proposal.Group, msgs, nil
}
