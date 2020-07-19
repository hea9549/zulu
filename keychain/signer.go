package keychain

type Signer interface {
	ID() string
	Sign(payload []byte) (Signature, error)
}
