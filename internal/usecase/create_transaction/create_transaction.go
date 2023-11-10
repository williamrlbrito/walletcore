package createtransaction

import (
	"github.com/williamrlbrito/walletcore/internal/entity"
	"github.com/williamrlbrito/walletcore/internal/gateway"
)

type CreateTransactionInputDTO struct {
	AcountIDFrom string
	AcountIDTo   string
	Amount       float64
}

type CreateTransactionOutputDTO struct {
	ID string
}

type CreateTransactionUseCase struct {
	TransactionGateway gateway.TransactionGateway
	AcountGateway      gateway.AccountGateway
}

func NewCreateTransactionUseCase(transactionGateway gateway.TransactionGateway, acountGateway gateway.AccountGateway) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		TransactionGateway: transactionGateway,
		AcountGateway:      acountGateway,
	}
}

func (useCase *CreateTransactionUseCase) Execute(input CreateTransactionInputDTO) (*CreateTransactionOutputDTO, error) {
	accountFrom, err := useCase.AcountGateway.FindById(input.AcountIDFrom)
	if err != nil {
		return nil, err
	}

	accountTo, err := useCase.AcountGateway.FindById(input.AcountIDTo)
	if err != nil {
		return nil, err
	}

	transaction, err := entity.NewTransaction(accountFrom, accountTo, input.Amount)

	if err != nil {
		return nil, err
	}

	err = useCase.TransactionGateway.Create(transaction)

	if err != nil {
		return nil, err
	}

	return &CreateTransactionOutputDTO{ID: transaction.ID}, nil
}
