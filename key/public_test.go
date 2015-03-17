package key

import (
	"btc/key/ec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)


var _ = Describe("Public", func() {
	Describe("#Hex", func() {
		It("has the correct value", func() {
			//generator point
			 x := ec.W("55066263022277343669578718895168534326250603453777594175500187360389116729240")  //0x79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798
			 y := ec.W("32670510020758816978083085130507043184471273380659243275938904335757337482424") // 0x483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b8

			public := &Public{&ec.Point{x,y}}
			Expect(public.Hex()).To(Equal("02483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b8"))
			Expect(public.FullString()).To(Equal("0479be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b8"))

		})
	})
})


