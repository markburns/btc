package keys

import (
	"math/big"
	"encoding/hex"
)


type PrivateKey struct{
  Key *big.Int
}

func(k *PrivateKey) Hex() string{
 return hex.EncodeToString(k.Key.Bytes())
}

func (k *PrivateKey) String() string{
	return k.Hex()
	
}
