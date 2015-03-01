package keys_test

import (
	. "btc/keys"
	"btc/keys/ec"
	"math/big"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type MockRandomNumberGenerator struct { }

func(m *MockRandomNumberGenerator) Random() *big.Int{
	fairDiceRoll := "4"
	i, _ := big.NewInt(0).SetString(fairDiceRoll, 10)
	return i
}

var _ = Describe("KeyGenerator", func() {
	Describe("#NewKeyPair", func() {
		It("generates a new public and private key", func() {
			random    := &MockRandomNumberGenerator{}
			curve     := ec.NewCurve()
			generator := &KeyGenerator{random, curve}
			pair      := generator.NewKeyPair()

			Expect(pair.PublicKey .String()).To(Equal("1234"))
			Expect(pair.PrivateKey.String()).To(Equal("1234"))
		})
	})
})
