package keys

import (
	"math/big"
	. "btc/keys/ec"
)

type RandomNumberGenerator interface {
	Random() *big.Int
}

type KeyGenerator struct {
	RandomNumberGenerator
	*Curve
}

func (k *KeyGenerator) NewKeyPair() *KeyPair {
	return k.calculate()
}

func (k *KeyGenerator) calculate() *KeyPair{
	private := B("1234")
	public  := B("1234")
	return &KeyPair{private, public}
}

