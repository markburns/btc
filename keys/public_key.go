package keys

import (
	"btc/keys/ec"
	"encoding/hex"
)

type PublicKey struct{
  key *ec.Point
}

func(k *PublicKey) Hex() string{
	return k.toHex(k.key.X.Bytes())
}

func (k *PublicKey) String() string{
	hex := k.Hex()

	if(ec.Odd(k.key.Y)){
		return "03" + hex
	}

	return "02" + hex
}

func (k *PublicKey) FullString() string{
	return "04" + k.toHex(k.key.X.Bytes()) + k.toHex(k.key.Y.Bytes())
}

func(k *PublicKey) toHex(s []byte) string{
	return hex.EncodeToString(s)
}
