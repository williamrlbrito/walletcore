package create_account

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/williamrlbrito/walletcore/internal/entity"
	"github.com/williamrlbrito/walletcore/internal/usecase/mocks"
)

func TestCreateAccountUseCase(t *testing.T) {
	client, _ := entity.NewClient("Jhon Doe", "j@j")
	clientGatewayMock := &mocks.ClientGatewayMock{}
	clientGatewayMock.On("Get", mock.Anything).Return(client, nil)

	accountGatewayMock := &mocks.AccountGatewayMock{}
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
