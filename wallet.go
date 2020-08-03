package main

import "github.com/DE-labtory/zulu/db/leveldb"

type WalletId string

type Wallet struct {
	Id            WalletId
	signerPayload []byte
}

func NewWallet(password string, payload interface{}) *Wallet {
	// using password, payload ( random io or some byte seed or etc. )
	// encrypt payload using password and return new wallet objects ( using AES, etc. )
	panic("impl me!")
}
func (w *Wallet) GetSignerPayload(password string) ([]byte, error) {
	panic("impl me!")
}

type WalletStore interface {
	Get(id WalletId) *Wallet
	Set(wallet Wallet) error
}

// todo split below impl code to other file
type LevelDBWalletStore struct {
	db leveldb.DBProvider
}

func (l *LevelDBWalletStore) Get(id WalletId) *Wallet {
	panic("implement me")
}

func (l *LevelDBWalletStore) Set(wallet Wallet) error {
	panic("implement me")
}

func NewLevelDBWalletStore(path string) (*LevelDBWalletStore, error) {
	panic("impl me!")
}
