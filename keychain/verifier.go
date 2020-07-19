package keychain

type Verifier interface {
	ID() string
	Verify(hash []byte, data interface{}) (bool, error)
}
