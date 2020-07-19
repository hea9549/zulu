package keychain

type Generator interface {
	Gen(payload interface{}) (Signer, Verifier, error)
}
