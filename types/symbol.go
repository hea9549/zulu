package types

type Symbol string

const (
	BTC  = Symbol("btc")
	ETH  = Symbol("eth")
	USDT = Symbol("usdt")
	KLAY = Symbol("klay")
)

func (s Symbol) Platform() Platform {
	switch s {
	case BTC:
		return Bitcoin
	case ETH:
	case USDT:
		return Ethereum
	case KLAY:
		return Klaytn
	default:
		break
	}
	return Unknown
}
