package keys

import (
	"btc/keys/ec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)


var _ = Describe("Key", func() {
	Describe("#Hex", func() {
		It("has the correct value", func() {
			value := ec.B("72759466100064397073952777052424474334519735946222029294952053344302920927294")
			private := &PrivateKey{value}
			Expect(private.Hex()).To(Equal("a0dc65ffca799873cbea0ac274015b9526505daaaed385155425f7337704883e"))

			x := ec.B("3423904187495496827825042940737875085827330420143621346629173781207857376010")
			y := ec.B("75711134420273723792089656449854389054866833762486990555172221523628676983696")
			public := &PublicKey{&ec.Point{x,y}}
			Expect(public.Hex()).To(Equal("020791dc70b75aa995213244ad3f4886d74d61ccd3ef658243fcad14c9ccee2b0a"))
			Expect(public.FullString()).To(Equal("040791dc70b75aa995213244ad3f4886d74d61ccd3ef658243fcad14c9ccee2b0aa762fbc6ac0921b8f17025bb8458b92794ae87a133894d70d7995fc0b6b5ab90"))
		})
	})
})


