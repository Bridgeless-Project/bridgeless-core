package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParamsValidate(t *testing.T) {
	validAdmin := sdk.AccAddress(make([]byte, 20)).String()

	require.NoError(t, NewParams(validAdmin).Validate())
	require.NoError(t, DefaultParams().Validate())
	require.Error(t, NewParams("invalid").Validate())
}
