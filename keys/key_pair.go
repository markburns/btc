package keys

import (
	"math/big"
	. "btc/keys/ec"
)

type RandomNumber interface {
	Random() *big.Int
}

type KeyPair struct {
	publicKey *big.Int
	privateKey *big.Int
	calculated bool
}

func NewKeyPair(r RandomNumber) *KeyPair {
	return &KeyPair{B("1234"), B("1234"), false}
}

func(k *KeyPair) PublicKey() *big.Int {
	k.calculate() 

	return k.publicKey
}

func(k *KeyPair) PrivateKey() *big.Int {
	k.calculate() 

	return k.privateKey
}

func (k *KeyPair) calculate() {
	if(!k.calculated){ 
		k.privateKey = B("1234")
		k.publicKey = B("1234")
		k.calculated = true
}
}

