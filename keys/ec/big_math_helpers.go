package ec

import (
	"math/big"
)

func exp(a, b, p *big.Int) *big.Int{
	return calc().Exp(a, b, p)
}
func mul(a, b *big.Int) *big.Int{
	return calc().Mul(a, b)
}
func sub(a, b *big.Int) *big.Int{
	return calc().Sub(a, b)
}

func DivMod(a, b, p *big.Int) *big.Int{
	a, b, p = dup(a), dup(b), dup(p)
	return divMod(a,b,p)
}

func divMod(a, b, p *big.Int) *big.Int{
	r, _ := calc().DivMod(a, b, calc())
	
	return mod(r, p)
}

func mod(a, p *big.Int) *big.Int{
	return calc().Mod(a, p)
}

func add(a, b *big.Int) *big.Int{
	return calc().Add(a, b)
}

func calc() * big.Int{
	return B("0")
}

func eq(a *big.Int, b int) bool{
	return cmp(a, b) == 0
}
func Eq(a *big.Int, b int) bool{
	return eq(a, b)
}
func Lt(a *big.Int, b int) bool{
	return lt(a, b)
}

func lt(a *big.Int, b int) bool{
	return cmp(a, b) < 0
}

func Gt(a *big.Int, b int) bool{
	return gt(a, b)
}
func gt(a *big.Int, b int) bool{
	return cmp(a, b) > 0
}
func gte(a *big.Int, b int) bool{
	return gt(a,b) || eq(a,b)
}

func cmp(a *big.Int, b int) int{
	return a.Cmp(big.NewInt(int64(b)))
}

func dup(x *big.Int) * big.Int{
	return mul(B("1"), x)
}

func B(s string) * big.Int{
	x, _ := big.NewInt(0).SetString(s, 10)

	return x
}
