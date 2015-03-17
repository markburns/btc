package key

import (
	. "btc/key/ec"
	"btc/random"
)

type Generator struct {
	rn random.Random
	curve *Curve
}

func NewGenerator() *Generator {
	random := random.System{}
	return &Generator{random, NewCurve()}
}

func (k *Generator) NewPair() (*Private, *Public) {
	private := k.rn.Random()
	public  := k.curve.Multiply(private)

	return &Private{private}, &Public{public}
}




