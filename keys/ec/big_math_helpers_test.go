package ec_test

import (
	. "btc/keys/ec"

	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)


var _ = Describe("big math helper functions", func() {
	var TestDivMod = func(a,b,p,r string) {
		It(fmt.Sprintf("(%v / %v) mod %v = %v", a,b,p,r), func() {
			bigA, bigB, bigP := B(a), B(b), B(p)
			result := DivMod(bigA,bigB,bigP)

			//ensure it doesn't mutate the inputs
			Expect(bigA.String()).To(Equal(a))
			Expect(bigB.String()).To(Equal(b))
			Expect(bigP.String()).To(Equal(p))

			Expect(result.String()).To(Equal(r))
		})
	}

	Describe("#DivMod", func() {
		TestDivMod("1",  "1", "10", "1")
		TestDivMod("14", "2", "10", "7")
		TestDivMod("28", "2", "10", "4")
	})

	Describe("#cmp", func() {
		It("works", func() {
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
})


