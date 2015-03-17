package base58

import (
	"crypto/sha256"
	"golang.org/x/crypto/ripemd160"
)

func Hash160(b []byte) []byte {
	return Ripemd160(Sha256(b))
}

func DoubleSha(b []byte) []byte {
	return Sha256(Sha256(b))
}
func ReverseBytes(p []byte) []byte {
	for i, j := 0, len(p)-1; i < j; i, j = i+1, j-1 {
		p[i], p[j] = p[j], p[i]
	}
	return p
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
