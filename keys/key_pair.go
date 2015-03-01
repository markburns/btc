package keys

import (
	"math/big"
)

type RandomNumber interface {
	Random() *big.Int
}

type KeyPair struct {
	PublicKey *big.Int
	PrivateKey *big.Int
}
