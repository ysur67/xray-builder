package models

type KeyPair struct {
	Pub     string
	Private string
}

func NewKeyPair(pub string, private string) *KeyPair {
	return &KeyPair{
		Pub:     pub,
		Private: private,
	}
}
