package keychain

type Signature interface {
	Verify(hash []byte, verifier Verifier) (bool, error)
}
