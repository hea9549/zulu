package kms

import (
	"github.com/DE-labtory/zulu/keychain"
)

type SampleKMSSigner struct {
	id       string
	endpoint string
}

func (s SampleKMSSigner) Sign(payload []byte) (keychain.Signature, error) {
	// todo : request sign to kms server
	panic("impl me!")
}

func (s SampleKMSSigner) ID() string {
	return s.id
}

type SampleKMSVerifier struct {
	id       string
	endpoint string
}

func (s SampleKMSVerifier) Verify(hash []byte, data interface{}) (bool, error) {
	// todo : request verify to kms server using raw data
	panic("impl me!")
}
func (s SampleKMSVerifier) ID() string {
	return s.id
}

type SampleKMSSignature struct {
	payload []byte
}

func (s SampleKMSSignature) Verify(hash []byte, verifier keychain.Verifier) (bool, error) {
	return verifier.Verify(hash, s.payload)
}

type SampleKMSGenerator struct {
}

func (s SampleKMSGenerator) Gen(payload interface{}) (keychain.Signer, keychain.Verifier, error) {
	// todo : set kms endpoint using payload and return signing, verifying object.
	panic("implement me")
}
