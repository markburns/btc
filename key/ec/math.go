package ec

import (
	"math/big"
)

var calc *big.Int

func init(){
	calc = big.NewInt(0)
}

func Eq(a *Wif, b int) bool{
	return eq(a, b)
}

func Lt(a *Wif, b int) bool{
	return lt(a, b)
}

func Gt(a *Wif, b int) bool{
	return gt(a, b)
}

func Odd(x *Wif) bool{
	return eq(mod(x, W("2")), 1)
}

func exp(a, b, p *Wif) *Wif{
	return W(calc.Exp(a.BigInt(), b.BigInt(), p.BigInt()))
}

func add(a, b *Wif) *Wif{
	return W(calc.Add(a.BigInt(), b.BigInt()))
}

func sub(a, b *Wif) *Wif{
	return W(calc.Sub(a.BigInt(), b.BigInt()))
}

func mul(a, b *Wif) *Wif{
	return W(calc.Mul(a.BigInt(), b.BigInt()))
}

func div(a, b *Wif) *Wif{
	r, _ := calc.DivMod(a.BigInt(), b.BigInt(), calc)

	return W(r)
}

func modinv(a, b *Wif) *Wif{
	return W(a.ModInverse(a.BigInt(), b.BigInt()))
}

func mod(a, b *Wif) *Wif{
	return W(calc.Mod(a.BigInt(), b.BigInt()))
}

func lt(a *Wif, b int) bool{
	return cmp(a, b) < 0
}

func eq(a *Wif, b int) bool{
	return cmp(a, b) == 0
}

func gte(a *Wif, b int) bool{
	return gt(a,b) || eq(a,b)
}

func gt(a *Wif, b int) bool{
	return cmp(a, b) > 0
}

func cmp(a *Wif, b int) int{
	return a.Cmp(big.NewInt(int64(b)))
}

func dup(x *Wif) * Wif{
	return mul(W("1"), x)
}
