package key

import (
	"btc/key/ec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)


var _ = Describe("Private", func() {
	Describe("#Hex", func() {
		It("has the correct value", func() {
			value := ec.W("72759466100064397073952777052424474334519735946222029294952053344302920927294")

			private := &Private{value}
			Expect(private.Hex()).To(Equal("00a0dc65ffca799873cbea0ac274015b9526505daaaed385155425f7337704883e"))
		})
	})
})


