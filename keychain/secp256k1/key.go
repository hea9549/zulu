package secp256k1

import (
	"crypto/ecdsa"
	"errors"
	"github.com/DE-labtory/zulu/keychain"
	"github.com/btcsuite/btcd/btcec"
	"io"
	"math/big"
)

var curve = btcec.S256()
var curveHalfOrder = new(big.Int).Rsh(curve.N, 1)

type PublicKey struct {
	id  string
	key *ecdsa.PublicKey
}

func (p PublicKey) ID() string {
	return p.id
}

func (p PublicKey) Verify(hash []byte, data interface{}) (bool, error) {
	panic("implement me")
}

type PrivateKey struct {
	id  string
	key *ecdsa.PrivateKey
}

func NewSecp256k1Key(id string, rand io.Reader) (*PrivateKey, error) {
	keyData, err := ecdsa.GenerateKey(btcec.S256(), rand)
	if err != nil {
		return nil, err
	}

	prvKey := &PrivateKey{
		id:  id,
		key: keyData,
	}
	return prvKey, nil
}

// This is borrowed from "github.com/btcsuite/btcd/btcec"
func (p PrivateKey) Sign(hashPayload []byte) (keychain.Signature, error) {
	privkey := p.key
	N := curve.N
	halfOrder := curveHalfOrder
	k := nonceRFC6979(privkey.D, hashPayload)
	inv := new(big.Int).ModInverse(k, N)
	r, _ := privkey.Curve.ScalarBaseMult(k.Bytes())
	r.Mod(r, N)

	if r.Sign() == 0 {
		return nil, errors.New("calculated R is zero")
	}

	e := hashToInt(hashPayload, privkey.Curve)
	s := new(big.Int).Mul(privkey.D, r)
	s.Add(s, e)
	s.Mul(s, inv)
	s.Mod(s, N)

	if s.Cmp(halfOrder) == 1 {
		s.Sub(N, s)
	}
	if s.Sign() == 0 {
		return nil, errors.New("calculated S is zero")
	}

	return &Signature{R: r, S: s}, nil
}
func (p PrivateKey) ID() string {
	return p.id
}

func (p PrivateKey) FromByte(rawData []byte) error {
	panic("implement me")
}

func (p PrivateKey) PublicKey() PublicKey {
	return PublicKey{
		id:  p.id,
		key: &p.key.PublicKey,
	}
}
