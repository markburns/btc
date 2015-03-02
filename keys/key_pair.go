package keys

type KeyPair struct {
	privateKey *PrivateKey
	publicKey *PublicKey
}

func (k *KeyPair) PrivateKey() *PrivateKey {
	return k.privateKey
}

func (k *KeyPair) PublicKey() *PublicKey {
	return k.publicKey
}
