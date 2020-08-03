package coin

import (
	"github.com/DE-labtory/zulu/signer"
	"github.com/DE-labtory/zulu/types"
)

type Service interface {
	Transfer(account types.Account, to types.Address, signer signer.Signer) ([]byte, error)
	CreateAccount(signer signer.Signer) (types.Account, error)
}
