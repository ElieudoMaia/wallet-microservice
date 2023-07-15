package gateway

import "github.com/elieudomaia/ms-wallet-app/internal/entity"

type ClientGateway interface {
	Get(id string) (*entity.Client, error)
	Save(client *entity.Client) error
}
