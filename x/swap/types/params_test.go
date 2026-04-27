package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestParamsValidate(t *testing.T) {
	validAdmin := sdk.AccAddress(make([]byte, 20)).String()
	validWrappedBridge := "0x2222222222222222222222222222222222222222"
	validSwapper := "0x1111111111111111111111111111111111111111"

	require.NoError(t, NewParams(validAdmin, validWrappedBridge, validSwapper, DefaultSwapDeadlineSeconds).Validate())
	require.NoError(t, DefaultParams().Validate())
	require.Error(t, NewParams("invalid", validWrappedBridge, validSwapper, DefaultSwapDeadlineSeconds).Validate())
	require.Error(t, NewParams(validAdmin, "invalid", validSwapper, DefaultSwapDeadlineSeconds).Validate())
	require.Error(t, NewParams(validAdmin, validWrappedBridge, "invalid", DefaultSwapDeadlineSeconds).Validate())
	require.Error(t, NewParams(validAdmin, validWrappedBridge, validSwapper, 0).Validate())
}
