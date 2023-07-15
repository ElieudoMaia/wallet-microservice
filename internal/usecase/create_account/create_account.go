package create_account

import (
	"errors"

	"github.com/elieudomaia/ms-wallet-app/internal/entity"
	"github.com/elieudomaia/ms-wallet-app/internal/gateway"
)

type CreateAccountInputDTO struct {
	ClientID string
}

type CreateAccountOutputDTO struct {
	ID string
}

type CreateAccountUseCase struct {
	accountGateway gateway.AccountGateway
	clientGateway  gateway.ClientGateway
}

func NewCreateAccountUseCase(accountGateway gateway.AccountGateway, clientGateway gateway.ClientGateway) *CreateAccountUseCase {
	return &CreateAccountUseCase{
		accountGateway: accountGateway,
		clientGateway:  clientGateway,
	}
}

func (uc *CreateAccountUseCase) Execute(input *CreateAccountInputDTO) (*CreateAccountOutputDTO, error) {
	client, err := uc.clientGateway.Get(input.ClientID)
	if err != nil {
		return nil, errors.New("client not found")
	}
	account, _ := entity.NewAccount(client)
	err = uc.accountGateway.Save(account)
	if err != nil {
		return nil, errors.New("error saving account")
	}
	return &CreateAccountOutputDTO{
		ID: account.ID,
	}, nil
}
