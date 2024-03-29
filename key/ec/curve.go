package ec

import (
	//"github.com/cznic/mathutil"
	"math/big"
)

type Curve struct {
	p *big.Int // prime
	a *big.Int // curve param
	b *big.Int // curve param
	g *Point   // generator point
	n *big.Int // number of points
	h *big.Int // cofactor
}

func NewCurve() *Curve {
	return NewSecp256k1()
}

// convert a string to *big.Int panics with an invalid value.
// Ideally this would exist only as a method in an internal package, but they
// don't exist in go yet. Hence this is public, but not intended for public use.
func W(s string) *big.Int {
	result, success := new(big.Int).SetString(s, 10)

	if !success {
		panic("Invalid value for big.Int")
	}
	return result
}

func (c *Curve) Prime() *big.Int {
	return dup(c.p)
}
func (c *Curve) GeneratorPoint() *Point {
	return &Point{dup(c.g.X), dup(c.g.Y)}
}

func (c *Curve) Mod(w *big.Int) *big.Int {
	return mod(w, c.p)
}

func (c *Curve) ModularInverse(a *big.Int) *big.Int {
	return modinv(a, c.p)
}

func (c *Curve) AddPoints(a, b *Point) *Point {
	lamAdd := mod(mul((sub(b.Y, a.Y)), c.ModularInverse(sub(b.X, a.X))), c.p)

	x := mod(
		sub(mul(lamAdd, lamAdd), add(a.X, b.X)),
		c.p,
	)

	y := mod(
		sub(
			mul(lamAdd, sub(a.X, x)),
			a.Y),
		c.p,
	)

	return &Point{x, y}
}

func (c *Curve) DoublePoint(a *Point) *Point {
	threeXSquaredPlusA := add(
		mul(
			W("3"),
			mul(a.X, a.X),
		),
		c.a,
	)

	modInvTwoY := c.ModularInverse(mul(W("2"), a.Y))

	lam := mod(mul(threeXSquaredPlusA, modInvTwoY), c.p)

	x := mod(
		sub(
			mul(lam, lam),
			mul(W("2"), a.X),
		),
		c.p,
	)

	y := mod(
		sub(
			mul(lam, (sub(a.X, x))),
			a.Y,
		),
		c.p,
	)

	return &Point{x, y}
}

func (c *Curve) Multiply(x *big.Int) *Point {
	q := c.GeneratorPoint()

	length := x.BitLen()

	for i := 0; i < length-1; i++ {
		bit := x.Bit(i)
		q = c.DoublePoint(q)

		if bit == 1 {
			q = c.AddPoints(q, c.GeneratorPoint())
		}
	}

	return q
}
