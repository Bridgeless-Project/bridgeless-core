package tx

import (
	errorsmod "cosmossdk.io/errors"
	bridgeTypes "github.com/Bridgeless-Project/bridgeless-core/v12/types"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetMsgs takes a slice of sdk.Msg's and turn them into Any's.
func SetMsgs(msgs []sdk.Msg) ([]*types.Any, error) {
	anys := make([]*types.Any, len(msgs))
	for i, msg := range msgs {
		var err error
		anys[i], err = types.NewAnyWithValue(msg)
		if err != nil {
			return nil, err
		}
	}
	return anys, nil
}

// GetMsgs takes a slice of Any's and turn them into sdk.Msg's.
func GetMsgs(anys []*types.Any, name string) ([]sdk.Msg, error) {
	msgs := make([]sdk.Msg, len(anys))
	for i, any := range anys {
		cached := any.GetCachedValue()
		if cached == nil {
			return nil, errorsmod.Wrapf(bridgeTypes.ErrFailedToGetMsgs, "any cached value is nil, %s messages must be correctly packed any values", name)
		}
		msgs[i] = cached.(sdk.Msg)
	}
	return msgs, nil
}

// UnpackInterfaces unpacks Any's to sdk.Msg's.
func UnpackInterfaces(unpacker types.AnyUnpacker, anys []*types.Any) error {
	for _, any := range anys {
		var msg sdk.Msg
		err := unpacker.UnpackAny(any, &msg)
		if err != nil {
			return err
		}
	}

	return nil
}
