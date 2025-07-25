package hd

import (
	"testing"

	"github.com/Bridgeless-Project/bridgeless-core/v12/types"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
)

func BenchmarkEthSecp256k1Algo_Derive(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		deriveFn := EthSecp256k1.Derive()
		if _, err := deriveFn(mnemonic, keyring.DefaultBIP39Passphrase, types.BIP44HDPath); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkEthSecp256k1Algo_Generate(b *testing.B) {
	bz, err := EthSecp256k1.Derive()(mnemonic, keyring.DefaultBIP39Passphrase, types.BIP44HDPath)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		(&ethSecp256k1Algo{}).Generate()(bz)
	}
}
