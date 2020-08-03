package eth

import (
	"github.com/DE-labtory/zulu/signer"
	"github.com/DE-labtory/zulu/types"
)

type Service struct {
}

func NewService() *Service {
	panic("impl me!")
}

func (s *Service) Transfer(account types.Account, to types.Address, signer signer.Signer) ([]byte, error) {
	panic("implement me")
}

func (s *Service) CreateAccount(signer signer.Signer) (types.Account, error) {
	panic("implement me")
}
