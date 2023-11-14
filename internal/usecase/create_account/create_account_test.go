package create_account

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/williamrlbrito/walletcore/internal/entity"
)

type ClientGatewayMock struct {
	mock.Mock
}

func (mock *ClientGatewayMock) Get(id string) (*entity.Client, error) {
	args := mock.Called(id)
	return args.Get(0).(*entity.Client), args.Error(1)
}

func (mock *ClientGatewayMock) Save(client *entity.Client) error {
	args := mock.Called(client)
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

func (mock *AccountGatewayMock) UpdateBalance(account *entity.Account) error {
	args := mock.Called(account)
	return args.Error(0)
}

func TestCreateAccountUseCase(t *testing.T) {
	client, _ := entity.NewClient("Jhon Doe", "j@j")
	clientGatewayMock := &ClientGatewayMock{}
	clientGatewayMock.On("Get", mock.Anything).Return(client, nil)

	accountGatewayMock := &AccountGatewayMock{}
	accountGatewayMock.On("Save", mock.Anything).Return(nil)

	useCase := NewCreateAccountUseCase(accountGatewayMock, clientGatewayMock)

	input := CreateAccountInputDTO{
		ClientId: client.ID,
	}

	output, err := useCase.Execute(input)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	clientGatewayMock.AssertExpectations(t)
	accountGatewayMock.AssertExpectations(t)
	clientGatewayMock.AssertNumberOfCalls(t, "Get", 1)
	accountGatewayMock.AssertNumberOfCalls(t, "Save", 1)
}
