package base58

import(
	"math"
	"math/big"
	"btc/key/ec"
	"fmt"
)
const alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

type Checker struct{
	value *big.Int
	version int
}

func Check(b *big.Int, version int) string{
	encoder := &Checker{b, version}
	return encoder.Check()
}

func p(msg string, b []byte){
	s := ec.ToHex(b)

	fmt.Println(msg, s)
}

func(a *Checker) Check() string{
	b := make([]byte, 0)
	b = append(b, byte(a.version))
	b = append(b, a.payload()...)
	b = append(b, a.Checksum()...)

	i := big.NewInt(0).SetBytes(b)
	bt := EncodeBig(i)
	return string(bt)
}

func EncodeBig(input *big.Int) []byte {
	out := make([]byte,0)
	n := big.NewInt(0)
	n.Set(input)
	radix := big.NewInt(58)
	zero := big.NewInt(0)

	for n.Cmp(zero) > 0 {
		mod := big.NewInt(0)
		n.DivMod(n, radix, mod)
		out = append(out, alphabet[mod.Int64()])
	}

	for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
		out[i], out[j] = out[j], out[i]
	}

	return out
}

func(a *Checker) Checksum() []byte{
	b := make([]byte, 0)
	b = append(b, byte(a.version))
	fmt.Println(byte(a.version))
	b = append(b, a.payload()...)

	b = DoubleSha(b)

	return (b)[:4]
}
func(a *Checker) IntChecksum() int{
	return calculateSum(a.Checksum())
}

func(a *Checker) payload() []byte{
	return a.value.Bytes()
}

func calculateSum(result []byte) int{
	sum := 0.0

	for i := 0; i < len(result); i++{
		a := float64(result[i])
		dec :=  a * math.Pow(256.0, float64(i))
		sum += dec
	}
	return int(sum)
}
