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
	public := ec.W("232778611149933427379623484630052427045471094092634266730616881027394583545770") // 0x0202a406624211f2abbdc68da3df929f938c3399dd79fac1b51b0e4ad1d26a47aa
	address := "14ueonBGyKHdAG4v6dihLEYfQRdjjSSgfTjnsw73AsjCz4Xvada"
	version := 0

	Describe("#Check", func() {
		It("encodes the address as base58 with checksum", func() {
			result := Check(public, version)

			Expect(result).To(Equal(address))
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
			checker := Checker{public, version}
			sum := checker.IntChecksum()

			Expect(sum).To(Equal(3048623411))
		})
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
