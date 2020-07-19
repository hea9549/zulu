package secp256k1

import (
	"errors"
	"github.com/DE-labtory/zulu/keychain"
	"math/big"
)

type Signature struct {
	R *big.Int
	S *big.Int
}

func (s Signature) Verify(hash []byte, verifier keychain.Verifier) (bool, error) {
	if v, ok := verifier.(PublicKey); ok {
		return v.Verify(hash, s)
	}
	return false, errors.New("unsupported verifier")
}
