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
			Expect(curve.p.String()).To(Equal("115792089237316195423570985008687907853269984665640564039457584007908834671663"))
			Expect(curve.a.String()).To(Equal("0"))
			Expect(curve.b.String()).To(Equal("7"))
			Expect(curve.g.X.String()).To(Equal("55066263022277343669578718895168534326250603453777594175500187360389116729240"))
			Expect(curve.g.Y.String()).To(Equal("32670510020758816978083085130507043184471273380659243275938904335757337482424"))

			Expect(curve.n.String()).To(Equal("115792089237316195423570985008687907852837564279074904382605163141518161494337"))
			Expect(curve.h.String()).To(Equal("1"))

		})
	})
})
