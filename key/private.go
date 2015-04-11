package key

import (
	"encoding/hex"
	"math/big"
	"btc/key/base58"
)

type Private struct{
  n *big.Int
}

func(p *Private) Hex() string{
	return hex.EncodeToString(p.n.Bytes())
}

func (p *Private) String() string{
	return p.Hex()
}

func (p *Private) Version() int{
	return 128
}

func(p *Private) Base58() string{
	return base58.Check(p.n.Bytes(), p.Version())
}
