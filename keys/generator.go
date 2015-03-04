package keys

import (
	. "btc/keys/ec"
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

func (k *Generator) NewPair() *Pair {
	private := k.rn.Random()
	public  := k.curve.Multiply(private)

	return &Pair{&Private{private}, &Public{public}}
}




