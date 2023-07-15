package gateway

import "github.com/elieudomaia/ms-wallet-app/internal/entity"

type AccountGateway interface {
	Save(account *entity.Account) error
	FindByID(id string) (*entity.Account, error)
}
