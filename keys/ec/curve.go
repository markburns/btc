package ec

import (
	"math/big"
)
type Point struct {
	X *big.Int
	Y *big.Int
}

type Curve struct {
	p *big.Int // prime
	a *big.Int // curve param
	b *big.Int // curve param
	g *Point   // generator point
	n *big.Int // number of points
	h *big.Int // cofactor
}

func NewCurve() *Curve{
	return NewSecp256k1()
}

func (c *Curve) Prime() *big.Int{
	return dup(c.p)
}
func (c *Curve) GeneratorPoint() *Point{
	return &Point{dup(c.g.X), dup(c.g.Y)}
}

func (c *Curve) Mod(a string) *big.Int{
	s, _ := big.NewInt(0).SetString(a, 10)
	return mod(s, c.p)
}

func (c *Curve) ModularInverse(a *big.Int) *big.Int{
	var ratio, nm, newNm *big.Int
	n := dup(c.p)

	lm,hm := B("1"), B("0")
	low, high  := mod(a, n), dup(n)

	for gt(low,1){
		ratio = divMod(high, low, n)

		nm, newNm = sub(hm, mul(lm, ratio)), sub(high, mul(low, ratio))
		lm, low, hm, high = dup(nm), dup(newNm), dup(lm), dup(low)
	}

	return mod(lm, n)
}

func (c *Curve) AddPoints(a, b *Point) *Point{
	lam_add := mod(mul((sub(b.Y, a.Y )), c.ModularInverse(sub(b.X, a.X))), c.p)

	x := mod(
		sub(mul(lam_add,lam_add), add(a.X,b.X)),
		c.p,
	)

	y := mod(
		sub(
			mul(lam_add, sub(a.X,x)),
			a.Y),
		c.p,
	)

	return &Point{x, y}
}

func (c *Curve) DoublePoint(a *Point) *Point{
	three_x_squared_plus_a := add(
		mul(
			B("3"),
			mul(a.X, a.X),
		),
		c.a,
	)

	mod_inv_two_y := c.ModularInverse(mul(B("2"), a.Y))

	lam := mod(mul(three_x_squared_plus_a, mod_inv_two_y), c.p)

	x := mod(
		sub(
			mul(lam, lam),
			mul(B("2"), a.X),
		),
		c.p,
	)

	y := mod(
		sub(
			mul(lam, (sub(a.X,x))),
			a.Y,
		),
		c.p,
	)

	return &Point{x,y}
}

func (c *Curve) Multiply(x *big.Int) *Point{
	q := c.GeneratorPoint()

	length := x.BitLen()

	for i:= 0; i < length-1; i++{
		bit := x.Bit(i)
		q = c.DoublePoint(q)

		if(bit == 1){
		  q = c.AddPoints(q,c.GeneratorPoint())
		}
	}

	return q
}
