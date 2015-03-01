package keys_test

import (
	. "btc/keys"

	"math/big"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type MockRandomNumberGenerator struct {
}

func(m *MockRandomNumberGenerator) Random() *big.Int{
	fairDiceRoll := "4"
	i, _ := big.NewInt(0).SetString(fairDiceRoll, 10)
	return i
}


var _ = Describe("btc", func() {
	Describe("#NewKeyPair", func() {
		It("generates a new public and private key", func() {
			pair := NewKeyPair(&MockRandomNumberGenerator{})

			Expect(pair.PublicKey() .String()).To(Equal("1234"))
			Expect(pair.PrivateKey().String()).To(Equal("1234"))
		})
	})
})
