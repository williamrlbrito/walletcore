package create_client

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/williamrlbrito/walletcore/internal/usecase/mocks"
)

func TestCreateClientUseCase(t *testing.T) {
	clientGatewayMock := &mocks.ClientGatewayMock{}
	clientGatewayMock.On("Save", mock.Anything).Return(nil)

	useCase := NewCreateClientUseCase(clientGatewayMock)

	output, err := useCase.Execute(CreateClientInputDTO{
		Name:  "Jhon Doe",
		Email: "j@j",
	})

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ID)
	assert.Equal(t, "Jhon Doe", output.Name)
	assert.Equal(t, "j@j", output.Email)
	clientGatewayMock.AssertExpectations(t)
	clientGatewayMock.AssertNumberOfCalls(t, "Save", 1)
}
