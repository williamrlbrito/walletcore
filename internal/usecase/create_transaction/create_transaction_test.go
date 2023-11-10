package createtransaction

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/williamrlbrito/walletcore/internal/entity"
)

type TransactionGatewayMock struct {
	mock.Mock
}

func (mock *TransactionGatewayMock) Create(transaction *entity.Transaction) error {
	args := mock.Called(transaction)
	return args.Error(0)
}

type AccountGatewayMock struct {
	mock.Mock
}

func (mock *AccountGatewayMock) Save(account *entity.Account) error {
	args := mock.Called(account)
	return args.Error(0)
}

func (mock *AccountGatewayMock) FindById(id string) (*entity.Account, error) {
	args := mock.Called(id)
	return args.Get(0).(*entity.Account), args.Error(1)
}

func TestCreateTransactionUseCase(t *testing.T) {
	clientFrom, _ := entity.NewClient("Jhon Doe", "john@doe.com")
	accountFrom := entity.NewAccount(clientFrom)
	accountFrom.Credit(1000)

	clientTo, _ := entity.NewClient("Jane Doe", "jane@doe.com")
	accountTo := entity.NewAccount(clientTo)
	accountTo.Credit(1000)

	mockAccount := &AccountGatewayMock{}
	mockAccount.On("FindById", accountFrom.ID).Return(accountFrom, nil)
	mockAccount.On("FindById", accountTo.ID).Return(accountTo, nil)

	mockTransaction := &TransactionGatewayMock{}
	mockTransaction.On("Create", mock.Anything).Return(nil)

	input := CreateTransactionInputDTO{
		AcountIDFrom: accountFrom.ID,
		AcountIDTo:   accountTo.ID,
		Amount:       100,
	}

	useCase := NewCreateTransactionUseCase(mockTransaction, mockAccount)

	output, err := useCase.Execute(input)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	mockAccount.AssertExpectations(t)
	mockTransaction.AssertExpectations(t)
	mockAccount.AssertNumberOfCalls(t, "FindById", 2)
	mockTransaction.AssertNumberOfCalls(t, "Create", 1)
}
