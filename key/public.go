package key

import (
	. "btc/key/ec"
	"btc/key/base58"
	"encoding/hex"
	"math/big"
)

type Public struct{
  key *Point
}

func(p *Public) Hex() string{
	hex := p.toHex(p.y().Bytes())

	if(Odd(W(p.y()))){
		return "03" + hex
	}

	return "02" + hex
}

func(p *Public) BigInt() *big.Int{
	return p.y()
}

func(p *Public) Version() int{
	return 0
}

func(p *Public) Base58() string{
	return base58.Check(p.key.Y, p.Version())
}

func(p *Public) y() *big.Int{
	return p.key.Y.BigInt()
}

func(p *Public) x() *big.Int{
	return p.key.X.BigInt()
}

func (p *Public) String() string{
	return p.Hex()
}

func (p *Public) FullString() string{
	return "04" + p.toHex(p.x().Bytes()) + p.toHex(p.y().Bytes())
}

func(p *Public) toHex(b []byte) string{
	return hex.EncodeToString(b)
}
