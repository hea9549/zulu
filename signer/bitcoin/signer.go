package bitcoin

type Signer struct {
}

func newSigner(payload interface{}) *Signer {
	// from bitcoin generator
	panic("impl me!")
}
func (s *Signer) Sign() interface{} {
	panic("implement me")
}

func (s *Signer) Public() interface{} {
	panic("implement me")
}
