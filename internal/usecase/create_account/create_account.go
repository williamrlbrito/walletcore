package create_account

import (
	"github.com/williamrlbrito/walletcore/internal/entity"
	"github.com/williamrlbrito/walletcore/internal/gateway"
)

type CreateAccountInputDTO struct {
	ClientId string
}

type CreateAccountOutputDTO struct {
	Id string
}

type CreateAccountUseCase struct {
	AccountGateway gateway.AccountGateway
	ClientGateway  gateway.ClientGateway
}

func NewCreateAccountUseCase(accountGateway gateway.AccountGateway, clientGateway gateway.ClientGateway) *CreateAccountUseCase {
	return &CreateAccountUseCase{
		AccountGateway: accountGateway,
		ClientGateway:  clientGateway,
	}
}

func (useCase *CreateAccountUseCase) Execute(input *CreateAccountInputDTO) (*CreateAccountOutputDTO, error) {
	client, err := useCase.ClientGateway.Get(input.ClientId)
	if err != nil {
		return nil, err
	}

	account := entity.NewAccount(client)

	err = useCase.AccountGateway.Save(account)
	if err != nil {
		return nil, err
	}

	return &CreateAccountOutputDTO{
		Id: account.ID,
	}, nil
}
