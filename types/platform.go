package types

type Platform = string

const (
	Unknown  = Platform("unknown")
	Ethereum = Platform("ethereum")
	Bitcoin  = Platform("bitcoin")
	Klaytn   = Platform("klaytn")
)
