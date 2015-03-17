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
	fmt.Println(msg)

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
			result := fromHexString("01")
			x := W("0").SetBytes(result)
			Expect(x).To(Equal(W("1")))

		})
	})

	Describe("#PointFromHex", func() {
		It("generates a point", func() {
			a := PointFromHex("79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b8")
			Expect(a.X).To(Equal(W("55066263022277343669578718895168534326250603453777594175500187360389116729240")))
			Expect(a.Y).To(Equal(W("32670510020758816978083085130507043184471273380659243275938904335757337482424")))
		})

		It("satisfies the equation", func() {
			// bx ec-add 02c0dbc34ea407fd5457b4affe73751a33f351f24c2001ad4715a525d938a0fb9d 0000000000000000000000000000000000000000000000000000000000000000
			// 02c0dbc34ea407fd5457b4affe73751a33f351f24c2001ad4715a525d938a0fb9d
			a := PointFromHex("79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b8")

			satisfiesTheEquation("a", a)
		})
	})
})
