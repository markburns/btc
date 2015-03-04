package keys

import (
	"math/big"
	"encoding/hex"
)


type Private struct{
  n *big.Int
}

func(k *Private) Hex() string{
	return hex.EncodeToString(k.n.Bytes())
}

func (k *Private) String() string{
	return k.Hex()
	
}
