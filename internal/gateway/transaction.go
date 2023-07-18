package gateway

import "github.com/elieudomaia/ms-wallet-app/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
