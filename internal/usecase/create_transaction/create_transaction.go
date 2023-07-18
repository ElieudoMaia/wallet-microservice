package create_transaction

import (
	"errors"

	"github.com/elieudomaia/ms-wallet-app/internal/entity"
	"github.com/elieudomaia/ms-wallet-app/internal/gateway"
)

type CreateTransactionInputDTO struct {
	AccountIDFrom string
	AccountIDTo   string
	Amount        float64
}

type CreateTransactionOutputDTO struct {
	TransactionID string
}

type CreateTransactionUseCase struct {
	TransactionGateway gateway.TransactionGateway
	AccountGateway     gateway.AccountGateway
}

func NewCreateTransactionUseCase(transactionGateway gateway.TransactionGateway, accountGateway gateway.AccountGateway) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		TransactionGateway: transactionGateway,
		AccountGateway:     accountGateway,
	}
}

func (uc *CreateTransactionUseCase) Execute(input *CreateTransactionInputDTO) (*CreateTransactionOutputDTO, error) {
	accountFrom, err1 := uc.AccountGateway.FindByID(input.AccountIDFrom)
	if err1 != nil {
		return nil, errors.New("account from not found")
	}
	accountTo, err2 := uc.AccountGateway.FindByID(input.AccountIDTo)
	if err2 != nil {
		return nil, errors.New("account to not found")
	}
	transaction, err3 := entity.NewTransaction(accountFrom, accountTo, input.Amount)
	if err3 != nil {
		return nil, err3
	}
	err4 := uc.TransactionGateway.Create(transaction)
	if err4 != nil {
		return nil, err4
	}
	return &CreateTransactionOutputDTO{
		TransactionID: transaction.ID,
	}, nil
}
