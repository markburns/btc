package keys

import (
	"math/big"
	. "btc/keys/ec"
)

type RandomNumberGenerator interface {
	Random() *big.Int
}

type KeyGenerator struct {
	randomNumberGenerator RandomNumberGenerator
	curve *Curve
}

func (k *KeyGenerator) NewKeyPair() *KeyPair {
	private := k.randomNumberGenerator.Random()
	public  := k.curve.Multiply(private)

	return &KeyPair{&PrivateKey{private}, &PublicKey{public}}
}

