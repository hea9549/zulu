package ethereum

type Signer struct {
}

func newSigner(payload interface{}) *Signer {
	// from ethereum generator
	panic("impl me!")
}
func (s *Signer) Sign() interface{} {
	panic("implement me")
}

func (s *Signer) Public() interface{} {
	panic("implement me")
}
