package ec

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)


var _ = Describe("Curve", func() {
	Describe("#NewCurve", func() {
		It("sets up an instance with correct secp256k1 value", func() {
			curve := NewCurve()

			Expect(curve.Prime().String()).To(Equal("115792089237316195423570985008687907853269984665640564039457584007908834671663"))
		})
	})
})
