package ec

import (
	"math/big"
)

var calc *big.Int

func init() {
	calc = big.NewInt(0)
}

func Eq(a *big.Int, b int) bool {
	return eq(a, b)
}

func Lt(a *big.Int, b int) bool {
	return lt(a, b)
}

func Gt(a *big.Int, b int) bool {
	return gt(a, b)
}

func Odd(x *big.Int) bool {
	return eq(mod(x, W("2")), 1)
}

func exp(a, b, p *big.Int) *big.Int {
	return calc.Exp(a, b, p)
}

func add(a, b *big.Int) *big.Int {
	return calc.Add(a, b)
}

func sub(a, b *big.Int) *big.Int {
	return calc.Sub(a, b)
}

func mul(a, b *big.Int) *big.Int {
	return calc.Mul(a, b)
}

func div(a, b *big.Int) *big.Int {
	r, _ := calc.DivMod(a, b, calc)

	return r
}

func modinv(a, b *big.Int) *big.Int {
	return a.ModInverse(a, b)
}

func mod(a, b *big.Int) *big.Int {
	return calc.Mod(a, b)
}

func lt(a *big.Int, b int) bool {
	return cmp(a, b) < 0
}

func eq(a *big.Int, b int) bool {
	return cmp(a, b) == 0
}

func gte(a *big.Int, b int) bool {
	return gt(a, b) || eq(a, b)
}

func gt(a *big.Int, b int) bool {
	return cmp(a, b) > 0
}

func cmp(a *big.Int, b int) int {
	return a.Cmp(big.NewInt(int64(b)))
}

func dup(x *big.Int) *big.Int {
	return mul(W("1"), x)
}
