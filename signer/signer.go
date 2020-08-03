package signer

type Signer interface {
	Sign() interface{}   // can be a signature type (e.g. secp256k1, []byte, etc.)
	Public() interface{} // will return public data (e.g. secp256k1 pub point, []byte, etc.) with hex encoded string
}
