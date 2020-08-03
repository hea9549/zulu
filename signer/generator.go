package signer

type Payload []byte

type Generator interface {
	Gen(payload Payload) Signer
}
