package keychain

type PairStore interface {
	Save(s Signer, v Verifier) error
	Get(id string) (Signer, Verifier, error)
}

type KeyChain interface {
	GetSignerByID(id string) (Signer, error)
	GetVerifierByID(id string) (Verifier, error)
	GetPairByID(id string) (Signer, Verifier, error)
	CreateNewPair(generator Generator, payload interface{}) (Signer, Verifier, error)
}

type DefaultKeyChain struct {
	PairStore PairStore
}

func (d DefaultKeyChain) GetSignerByID(id string) (Signer, error) {
	s, _, err := d.PairStore.Get(id)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (d DefaultKeyChain) GetVerifierByID(id string) (Verifier, error) {
	_, v, err := d.PairStore.Get(id)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (d DefaultKeyChain) GetPairByID(id string) (Signer, Verifier, error) {
	s, v, err := d.PairStore.Get(id)
	if err != nil {
		return nil, nil, err
	}
	return s, v, nil
}

func (d DefaultKeyChain) CreateNewPair(generator Generator, payload interface{}) (Signer, Verifier, error) {
	s, v, err := generator.Gen(payload)
	if err != nil {
		return nil, nil, err
	}

	err = d.PairStore.Save(s, v)
	if err != nil {
		return nil, nil, err
	}

	return s, v, nil
}
