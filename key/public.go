package key

import (
	. "btc/key/ec"
	"btc/key/base58"
	"encoding/hex"

	"fmt"
)

type Public struct{
  key *Point
}

func(p *Public) Version() int{
	return 0
}

func(p *Public) String() string{
	return p.Base58()
}

func(p *Public) Base58() string{
	sha := base58.Hash160(p.HexCompressed())

	return base58.Check(sha, p.Version())
}

func(p *Public) Base58Full() string{
	sha := base58.Hash160(p.HexFull())
	result := base58.Check(sha, p.Version())
	fmt.Println(result)
	return result
}

func(p *Public) Hash160Full() string{
	return hex.EncodeToString(base58.Hash160(p.HexFull()))
}
func(p *Public) Hash160() string{
	return hex.EncodeToString(base58.Hash160(p.HexCompressed()))
}

func(p *Public) HexCompressedString() string{
 return hex.EncodeToString(p.HexCompressed())
}

func(p *Public) HexCompressed() []byte{
	hex := p.xBytes()

	var  value []byte
	if(Odd(p.key.Y)){
		value = append([]byte{0x03}, hex...)
	}

	value = append([]byte{0x02}, hex...)

	return value
}



func(p *Public) HexFullString() string{
 return hex.EncodeToString(p.HexFull())
}

func(p *Public) HexFull() []byte{
	b := make([]byte, 0)
	full := append(b,    p.xBytes()...)
	full =  append(full, p.yBytes()...)
	bytes := []byte{0x04}
	return append(bytes, full...)
}

func (p *Public) xBytes() []byte{
	return base58.DropLeadingZeroes(p.key.X.Bytes(), false)
}

func (p *Public) yBytes() []byte{
	return base58.DropLeadingZeroes(p.key.Y.Bytes(), false)
}
