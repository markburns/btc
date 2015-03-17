package key

import (
	"encoding/hex"
	"btc/key/ec"
	"btc/key/base58"
)

type Private struct{
  n *ec.Wif
}

func(p *Private) Hex() string{
	return "00" + hex.EncodeToString(p.n.Bytes())
}

func (p *Private) String() string{
	return p.Hex()
}

func (p *Private) Version() int{
	return 128
}

func(p *Private) Base58() string{
	return base58.Check(p.n, p.Version())
}
