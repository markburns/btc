package base58

//Address = Ripemd160(Sha256(PublicKey))
//public key
// "19YhkH2ZQHjcDQqMkMabtEEqY7ti4sv8es"
//private key
// "KybnQU4c2T65RzZbGn5Mq22HNH7GBoda2dhp1W6m1r94zVno2mpT"

import (
	"btc/key/ec"

	"encoding/hex"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)


var _ = Describe("base58", func() {
	Describe("#dropLeadingZeroes", func() {
		It("removes leading zeroes", func() {
			input    := []byte{byte(0),byte(0), byte(1), byte(0), byte(1)}
			result := DropLeadingZeroes(input, true)

			Expect(result[0]).To(Equal(byte(0)))
			Expect(result[1]).To(Equal(byte(1)))
			Expect(result[2]).To(Equal(byte(0)))
			Expect(result[3]).To(Equal(byte(1)))
		})
	})


	//0x3aba4162c7251c891207b747840551a71939b0de081f85c4e44cf7c13e41daa6
	private := ec.W("26563230048437957592232553826663696440606756685920117476832299673293013768870").Bytes()
	version := 128
	Describe("#Check", func() {
		It("encodes the address as base58 with checksum", func() {
			result := Check(private, version)

			expected := "5JG9hT3beGTJuUAmCQEmNaxAuMacCTfXuw1R3FCXig23RQHMr4K"
			Expect([]byte(result)).To(Equal([]byte(expected)))
		})
	})

	// bx base58check-decode
	// 14ueonBGyKHdAG4v6dihLEYfQRdjjSSgfTjnsw73AsjCz4Xvada
	// wrapper
	// {
	//     checksum 3048623411
	//     payload 0202a406624211f2abbdc68da3df929f938c3399dd79fac1b51b0e4ad1d26a47aa
	//     version 0
	// }

	Describe("#Checksum", func() {
		It("calculates the correct checksum", func() {
			checker := Checker{private, version}
			sum := checker.IntChecksum()

			Expect(sum).To(Equal(616040902))
		})
	})

	It("#encodeBig", func() {
		i := ec.W("239730442551110153596068370417601753742195052378289479559")
		result := string(encodeBig(i))
		Expect(result).To(Equal("thMirt546nngXqyPEz532S8fLwbozud8"))
	})

	It("#bigIntToBase58", func() {
		i := ec.W("239730442551110153596068370417601753742195052378289479559")
		b := []byte{0, 9, 198, 231, 17, 24, 216, 241, 43, 236, 107, 92, 97, 136, 75, 53, 103, 124, 10, 10, 227, 42, 2, 31, 135}

		result := bigIntToBase58(i,b)
		Expect(result).To(Equal("1thMirt546nngXqyPEz532S8fLwbozud8"))
	})
	Describe("#DoubleSha", func() {
		It("calculates the double sha", func() {
			// bx sha256 01
			// 4bf5122f344554c53bde2ebb8cd2b7e3d1600ad631c385a5d7cce23c7785459a
			// bx sha256 4bf5122f344554c53bde2ebb8cd2b7e3d1600ad631c385a5d7cce23c7785459a
			// 9c12cfdc04c74584d787ac3d23772132c18524bc7ab28dec4219b8fc5b425f70
			// => 70594042219340844655004043388548011718365913142946830322468519465099706130288

			bytes := DoubleSha([]byte{1})
			result := ec.W("0").SetBytes(bytes).String()
			expected := ec.W("70594042219340844655004043388548011718365913142946830322468519465099706130288").String()
			Expect(result).To(Equal(expected))
		})
	})
	Describe("#Sha256", func() {
		It("calculates the sha256", func() {
			bytes := Sha256([]byte{0})
			Expect(hex.EncodeToString(bytes)).To(Equal("6e340b9cffb37a989ca544e6bb780a2c78901d3fb33738768511a30617afa01d"))
		})
	})

})
