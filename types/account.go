package types

type Address string

type Account struct {
	Address Address `json:"address"`
	Coin    Coin    `json:"coin"`
	Balance Balance `json:"balance"`
}
