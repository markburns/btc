package key

import (
	"btc/key/ec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)


var _ = Describe("Public", func() {
	Describe("#Base58", func() {
		//Mastering bitcoin example 4-5
		x := ec.W("41637322786646325214887832269588396900663353932545912953362782457239403430124")
		y := ec.W("16388935128781238405526710466724741593761085120864331449066658622400339362166")
		public := &Public{&ec.Point{x, y}}

		It("#HexCompressed", func() {
			Expect(public.HexCompressedString()).   To(Equal("025c0de3b9c8ab18dd04e3511243ec2952002dbfadc864b9628910169d9b9b00ec"))
		})
	
		It("#Hash160", func() {
			Expect(public.Hash160Full()).   To(Equal("09c6e71118d8f12bec6b5c61884b35677c0a0ae3"))
		})


		It("#HexFull", func() {
			Expect(public.HexFullString()).To(Equal("045c0de3b9c8ab18dd04e3511243ec2952002dbfadc864b9628910169d9b9b00ec243bcefdd4347074d44bd7356d6a53c495737dd96295e2a9374bf5f02ebfc176"))
		})

		It("#Base58 (Compressed)", func() {
			Expect(public.Base58()).To(Equal("14cxpo3MBCYYWCgF74SWTdcmxipnGUsPw3"))
		})

		It("#Base58Full", func() {
			Expect(public.Base58Full()).To(Equal("1thMirt546nngXqyPEz532S8fLwbozud8"))
		})


		It("#String", func() {
			Expect(public.String()).          To(Equal("14cxpo3MBCYYWCgF74SWTdcmxipnGUsPw3"))
		})
	})


})


