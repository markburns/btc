package ec

import (
	"fmt"

	. "github.com/mgutz/ansi"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func p(s string) {
	fmt.Println(Color(s, "blue"))
}

func satisfiesTheEquation(msg string, c *Point) {
	curve := NewCurve()
	lhs := mod(mul(c.Y, c.Y), curve.p)
	rhs := mod(
		add(
			exp(c.X, W("3"), curve.p),
			W("7"),
		),
		curve.p,
	)

	Expect(lhs.String()).To(Equal(rhs.String()))

}

var _ = Describe("Point", func() {
	Describe("#splitHex", func() {
		It("generates a point", func() {
			a, b := splitHex("021bab84e687e36514eeaf5a017c30d32c1f59dd4ea6629da7970ca374513dd006")
			Expect(a).To(Equal("021bab84e687e36514eeaf5a017c30d32"))
			Expect(b).To(Equal("c1f59dd4ea6629da7970ca374513dd006"))
		})
	})

	Describe("#fromHexString", func() {
		It("creates a byte array correctly", func() {
			result := fromHexString("ff")
			x := W("0").SetBytes(result)
			Expect(x).To(Equal(W("255")))

		})
	})
})
