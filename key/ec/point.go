package ec

import (
	"encoding/hex"
	"fmt"
	"math/big"
)

type Point struct {
	X *big.Int
	Y *big.Int
}

func PointFromString(s string) *Point {
	s = s[2:]
	halfWay := len(s) / 2

	x := W(s[:halfWay])
	y := W(s[halfWay:])

	return &Point{X: x, Y: y}
}

func splitHex(s string) (a, b string) {
	halfWay := len(s) / 2

	a = s[0:halfWay]
	b = s[halfWay:len(s)]

	return a, b
}


func (p *Point) XHex() string {
	return ToHex(p.X.Bytes())
}

func (p *Point) YHex() string {
	return ToHex(p.Y.Bytes())
}

func fromHexString(s string) []byte {
	if len(s)%2 != 0 {
		s = "0" + s
	}

	b, err := hex.DecodeString(s)

	if err != nil {
		msg := fmt.Sprintf("error %v decoding string: %s", err, s)
		panic(msg)
	}
	return b
}

func ToHex(bytes []byte) string {
	b := make([]byte, len(bytes)*2)
	hex.Encode(b, bytes)

	return string(b)
}
