package ethereum

import (
	"github.com/DE-labtory/zulu/signer"
)

type SignerGenerator struct {
}

func NewSignerGenerator() *SignerGenerator {
	panic("impl me!")
}

func (s *SignerGenerator) Gen(payload signer.Payload) signer.Signer {
	/*
		ex, generate some signer using string(payload+"ETHEREUM") data as seed for signer
		that means, this gen method is derive signer method from root signer payload
	*/
	panic("implement me")
}
