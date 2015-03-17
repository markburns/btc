package ec

import (
	"github.com/cznic/mathutil"
)

type Curve struct {
	p *Wif // prime
	a *Wif // curve param
	b *Wif // curve param
	g *Point   // generator point
	n *Wif // number of points
	h *Wif // cofactor
}

func NewCurve() *Curve{
	return NewSecp256k1()
}

func (c *Curve) Prime() *Wif{
	return dup(c.p)
}
func (c *Curve) GeneratorPoint() *Point{
	return &Point{dup(c.g.X), dup(c.g.Y)}
}

func (c *Curve) Mod(w *Wif) *Wif{
	return mod(w, c.p)
}

func (c *Curve) ModularInverse(a *Wif) *Wif{
	return modinv(a, c.p)
}
func (c *Curve) YFrom(x *Wif) *Wif{
	//(y^2)% p ==(x^3 + 7)% p
	ySquared := add( exp(x, W("3"), c.p), W("7"))

	result := mod(ySquared, c.p).BigInt()
	return W(mathutil.SqrtBig(result))
}

func (c *Curve) AddPoints(a, b *Point) *Point{
	lamAdd := mod(mul((sub(b.Y, a.Y )), c.ModularInverse(sub(b.X, a.X))), c.p)

	x := mod(
		sub(mul(lamAdd,lamAdd), add(a.X,b.X)),
		c.p,
	)

	y := mod(
		sub(
			mul(lamAdd, sub(a.X,x)),
			a.Y),
			c.p,
		)

		return &Point{x, y}
	}

	func (c *Curve) DoublePoint(a *Point) *Point{
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
				mul(lam, (sub(a.X,x))),
				a.Y,
			),
			c.p,
		)

		return &Point{x,y}
	}

	func (c *Curve) Multiply(x *Wif) *Point{
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
