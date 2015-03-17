package random

import (
	"crypto/rand"
	"math/big"
	"btc/key/ec"
)

type Random interface {
	Random() *big.Int
}

type System struct {
}

func (s System) Random() (n *big.Int){
	curve := ec.NewSecp256k1()

	n, _ =  rand.Int(rand.Reader, curve.Prime().BigInt())

	return n
}
