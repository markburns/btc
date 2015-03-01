package ec_test

import (
	. "btc/keys/ec"

	"fmt"

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
	var TestDivMod = func(a,b,p,r string) {
		bigA, bigB, bigP := B(a), B(b), B(p)
		result := DivMod(bigA,bigB,bigP)
		Expect(bigA.String()).To(Equal(a))
		Expect(bigB.String()).To(Equal(b))
		Expect(bigP.String()).To(Equal(p))
		Expect(result.String()).To(Equal(r))
		return
		}

	Describe("#DivMod", func() {
		It("works", func() {
			TestDivMod("1", "1", "1", "1")
		})
	})

	Describe("#cmp", func() {
		It("works", func() {
			return
			Expect(Gt(B("1"), 0)).To(Equal(true))
			Expect(Gt(B("1123123123123891023801928309128309182309812309128309182093281039281321"), 5)).To(Equal(true))
			Expect(Gt(B("0"), 1)).To(Equal(false))
			Expect(Gt(B("0"), 0)).To(Equal(false))

			Expect(Lt(B("1"), 0)).To(Equal(false))
			Expect(Lt(B("1123123123123891023801928309128309182309812309128309182093281039281321"), 5)).To(Equal(false))
			Expect(Lt(B("0"), 1)).To(Equal(true))
			Expect(Lt(B("0"), 0)).To(Equal(false))

			Expect(Eq(B("1"), 0)).To(Equal(false))
			Expect(Eq(B("1123123123123891023801928309128309182309812309128309182093281039281321"), 5)).To(Equal(false))
			Expect(Eq(B("0"), 1)).To(Equal(false))
			Expect(Eq(B("0"), 0)).To(Equal(true))
			Expect(Eq(B("1"), 1)).To(Equal(true))



		})
	})

	Describe("#Y", func() {
		expectResult("0", "1")
		expectResult("1", "1")
		expectResult("2", "57896044618658097711785492504343953926634992332820282019728792003954417335832")

		expectResult(
			"103588105881134521297392950363556582805271693099531759694119316285034955766421",
			"45900550174020825671461307549774270965382867892451705830695672428573753728672",
		)
	})

	Describe("#AddPoints", func() {
		It("adds 2 points", func() {
			a := &Point{B("1"), B("2")}
			b := &Point{B("3"), B("4")}
			c := NewCurve()

			Expect(c.AddPoints(a, b).X.String()).To(Equal("115792089237316195423570985008687907853269984665640564039457584007908834671660"))
			Expect(c.AddPoints(a, b).Y.String()).To(Equal("2"))
		})
	})

	Describe("#DoublePoint", func() {
		It("double a point", func() {
			a := &Point{B("3"), B("5")}
			c := NewCurve()
			x := "19684655170343753222007067451476944335055897393158895886707789281344501894184"
			y := "51064311353656442181794804388831367363292063237547488741400794547487796090203"

			Expect(c.DoublePoint(a).X.String()).To(Equal(x))
			Expect(c.DoublePoint(a).Y.String()).To(Equal(y))
		})
	})

	Describe("#Multiply", func() {
		It("Multiplies the generator point by a value", func() {
			c := NewCurve()

			equiv := func(multiplier, x, y string){
				result := c.Multiply(B(multiplier))
				Expect(result.X.String()).To(Equal(x))
				Expect(result.Y.String()).To(Equal(y))
			}

			equiv("1", c.GeneratorPoint().X.String(), c.GeneratorPoint().Y.String())

			equiv("3",
				"112711660439710606056748659173929673102114977341539408544630613555209775888121",
				"25583027980570883691656905877401976406448868254816295069919888960541586679410",
			)

		})
	})


})

func expectResult(a, b string) {
	var title = fmt.Sprintf("the inverse of %s to be %s", a, b)

	It(title, func() {
		curve := NewCurve()
		r := curve.ModularInverse(B(a))
		Expect(r.String()).To(Equal(b))
	})
}
