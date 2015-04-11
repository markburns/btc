package base58

import (
	"crypto/sha256"
	"golang.org/x/crypto/ripemd160"
)

func Hash160(b []byte) []byte {
	result := Ripemd160(Sha256(b))
	return reverseBytes(DropLeadingZeroes(reverseBytes(result), false))
}

func DoubleSha(b []byte) []byte {
	return Sha256(Sha256(b))
}

func Sha256(b []byte) []byte {
	a := sha256.Sum256(b)

	return a[:32]
}

func Ripemd160(b []byte) []byte {
	algo := ripemd160.New()
	algo.Write(b)
	return algo.Sum(nil)[:len(b)]
}
