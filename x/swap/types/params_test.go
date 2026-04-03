package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestParamsValidate(t *testing.T) {
	validAdmin := sdk.AccAddress(make([]byte, 20)).String()
	validRouter := "0x1111111111111111111111111111111111111111"
	validWrappedBridge := "0x2222222222222222222222222222222222222222"

	require.NoError(t, NewParams(validAdmin, validRouter, validWrappedBridge, DefaultSwapDeadlineSeconds).Validate())
	require.NoError(t, DefaultParams().Validate())
	require.Error(t, NewParams("invalid", validRouter, validWrappedBridge, DefaultSwapDeadlineSeconds).Validate())
	require.Error(t, NewParams(validAdmin, "invalid", validWrappedBridge, DefaultSwapDeadlineSeconds).Validate())
	require.Error(t, NewParams(validAdmin, validRouter, "invalid", DefaultSwapDeadlineSeconds).Validate())
	require.Error(t, NewParams(validAdmin, validRouter, validWrappedBridge, 0).Validate())
}
