package main

import (
	"errors"
	"github.com/DE-labtory/zulu/coin"
	"github.com/DE-labtory/zulu/coin/btc"
	"github.com/DE-labtory/zulu/coin/eth"
	"github.com/DE-labtory/zulu/signer"
	"github.com/DE-labtory/zulu/signer/bitcoin"
	"github.com/DE-labtory/zulu/signer/ethereum"
	"github.com/DE-labtory/zulu/types"
)

var (
	ErrUnknownPlatform = errors.New("unknown platform error")
	ErrUnknownSymbol   = errors.New("unknown symbol error")
)

func Resolve(symbol types.Symbol) (coin.Service, signer.Generator, error) {
	service := getService(symbol)
	if service == nil {
		return nil, nil, ErrUnknownSymbol
	}

	platform := symbol.Platform()
	generator := getGenerator(platform)
	if generator == nil {
		return nil, nil, ErrUnknownPlatform
	}

	return service, generator, nil
}

func getService(symbol types.Symbol) coin.Service {
	switch symbol {
	case types.BTC:
		return btc.NewService()
	case types.ETH:
		return eth.NewService()
	case types.USDT:
		panic("impl me!")
	case types.KLAY:
		panic("impl me!")
	}
	return nil
}

func getGenerator(platform types.Platform) signer.Generator {
	switch platform {
	case types.Bitcoin:
		return bitcoin.NewSignerGenerator()
	case types.Ethereum:
		return ethereum.NewSignerGenerator()
	case types.Klaytn:
		panic("impl me!")
	}
	return nil
}
