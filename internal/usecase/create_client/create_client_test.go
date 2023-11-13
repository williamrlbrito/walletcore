package create_client

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

func TestCreateClientUseCase(t *testing.T) {
	clientGatewayMock := &ClientGatewayMock{}
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
