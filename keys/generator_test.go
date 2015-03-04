package keys

import (
	"btc/keys/ec"
	"math/big"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type MockRandomNumberGenerator struct { }
var dangerousDiceRoll = "72759466100064397073952777052424474334519735946222029294952053344302920927294"

func(m *MockRandomNumberGenerator) Random() *big.Int{
	return ec.B(dangerousDiceRoll)
}

var _ = Describe("Generator", func() {
	Describe("#NewPair", func() {
		var pair *Pair

		BeforeEach(func() {
			random    := &MockRandomNumberGenerator{}
			curve     := ec.NewCurve()

			generator := &Generator{random, curve}
			pair      = generator.NewPair()
		})

		It("generates a private key using the random number generator", func() {
			Expect(pair.Private().String()).To(Equal("72759466100064397073952777052424474334519735946222029294952053344302920927294"))
		})

		It("generates a public key using the random number generator", func() {
			Expect(pair.Public().String()).To(Equal("020791dc70b75aa995213244ad3f4886d74d61ccd3ef658243fcad14c9ccee2b0a"))
		})
	})
})
