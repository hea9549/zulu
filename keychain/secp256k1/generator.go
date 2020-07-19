package secp256k1

import (
	"crypto/rand"
	"errors"
	"github.com/DE-labtory/zulu/keychain"
	"github.com/google/uuid"
	"go/types"
	"io"
)

type Secp256k1 struct {
}

func (s Secp256k1) Gen(payload interface{}) (keychain.Signer, keychain.Verifier, error) {

	switch v := payload.(type) {
	case types.Nil:
		return genPairFromReader(rand.Reader)
	case io.Reader:
		return genPairFromReader(v)
	case []byte:
		pk := PrivateKey{}
		err := pk.FromByte(v)
		if err != nil {
			return nil, nil, err
		}
		return pk, pk.PublicKey(), nil

	default:
		break
	}
	return nil, nil, errors.New("unexpected payload type")
}

func genPairFromReader(reader io.Reader) (keychain.Signer, keychain.Verifier, error) {

	pk, err := NewSecp256k1Key(uuid.New().String(), reader)
	if err != nil {
		return nil, nil, err
	}
	return pk, pk.PublicKey(), nil
}
