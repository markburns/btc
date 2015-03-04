package keys

import (
	"btc/keys/ec"
	"encoding/hex"
)

type Public struct{
  key *ec.Point
}

func(k *Public) Hex() string{
	hex := k.toHex(k.key.Y.Bytes())

	if(ec.Odd(k.key.Y)){
		return "03" + hex
	}

	return "02" + hex
}

func (k *Public) String() string{
	return k.Hex()
}

func (k *Public) FullString() string{
	return "04" + k.toHex(k.key.X.Bytes()) + k.toHex(k.key.Y.Bytes())
}

func(k *Public) toHex(s []byte) string{
	return hex.EncodeToString(s)
}
