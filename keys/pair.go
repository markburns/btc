package keys

type Pair struct {
	private *Private
	public *Public
}

func (k *Pair) Private() *Private {
	return k.private
}

func (k *Pair) Public() *Public {
	return k.public
}

