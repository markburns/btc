package base58

import (
	"math"
	"math/big"

	//"fmt"
)

const alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

type Checker struct {
	payload []byte
	version int
}

func Check(b []byte, version int) string {
	encoder := &Checker{b, version}
	return encoder.Check()
}

func (a *Checker) Check() string {
	b := make([]byte, 0)
	b = append(b, byte(a.version))
	b = append(b, a.payload...)
	b = append(b, a.Checksum()...)

	i := big.NewInt(0).SetBytes(b)
	return bigIntToBase58(i, b)
}

func DropLeadingZeroes(b []byte, pad bool) []byte {
	output := []byte{}
	leadingZero := true
	index := 0

	for i := 0; i < len(b); i++{
		if(!leadingZero){
			output = append(output, b[i])
			index += 1
		}else{
			if(byte(b[i]) != byte(0)){
				output = append(output, b[i])
				leadingZero = false
			}
		}
	}
	if pad{
		if(len(output) % 2 == 1){
			return append([]byte{0}, output...)
		}
	}

	return output
}

func reverseBytes(b []byte) []byte {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return b
}

func encodeBig(input *big.Int) []byte {
	out := make([]byte, 0)
	n := big.NewInt(0)
	n.Set(input)
	radix := big.NewInt(58)
	zero := big.NewInt(0)

	for n.Cmp(zero) > 0 {
		mod := big.NewInt(0)
		n.DivMod(n, radix, mod)
		out = append(out, alphabet[mod.Int64()])
	}

	return reverseBytes(out)
}

//from https://github.com/ThePiachu/Split-Vanity-Miner-Golang/blob/master/src/pkg/mymath/base58.go
//MIT licence https://github.com/ThePiachu/Split-Vanity-Miner-Golang/blob/master/LICENSE
func bigIntToBase58(input *big.Int, val []byte) string{
	tmp:= encodeBig(input)//encoding of the number without zeroes in front
	
	//looking for zeros at the beginning
	i:=0
	for i=0; val[i]==0 && i<len(val); i++{
	}

	answer:=""
	for j:=0 ; j < i; j++{
		answer += "1"
	}
	answer+=string(tmp)//concatenates

	return answer
}

func (a *Checker) Checksum() []byte {
	b := make([]byte, 0)
	b = append(b, byte(a.version))
	b = append(b, DropLeadingZeroes(a.payload, false)...)

	b = DoubleSha(b)

	return (b)[:4]
}
func (a *Checker) IntChecksum() int {
	return calculateSum(a.Checksum())
}

func calculateSum(result []byte) int {
	sum := 0.0

	for i := 0; i < len(result); i++ {
		a := float64(result[i])
		dec := a * math.Pow(256.0, float64(i))
		sum += dec
	}
	return int(sum)
}
