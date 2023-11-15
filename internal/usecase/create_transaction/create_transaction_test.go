package create_transaction

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/williamrlbrito/walletcore/internal/entity"
	"github.com/williamrlbrito/walletcore/internal/event"
	"github.com/williamrlbrito/walletcore/internal/usecase/mocks"
	"github.com/williamrlbrito/walletcore/pkg/events"
)

func TestCreateTransactionUseCase(t *testing.T) {
	clientFrom, _ := entity.NewClient("Jhon Doe", "john@doe.com")
	accountFrom := entity.NewAccount(clientFrom)
	accountFrom.Credit(1000)

	clientTo, _ := entity.NewClient("Jane Doe", "jane@doe.com")
	accountTo := entity.NewAccount(clientTo)
	accountTo.Credit(1000)

	mockUow := &mocks.UowMock{}
	mockUow.On("Do", mock.Anything, mock.Anything).Return(nil)

	input := CreateTransactionInputDTO{
		AccountIDFrom: accountFrom.ID,
		AccountIDTo:   accountTo.ID,
		Amount:        100,
	}

	dispatcher := events.NewEventDispatcher()
	event := event.NewTransactionCreatedEvent()
	ctx := context.Background()

	useCase := NewCreateTransactionUseCase(mockUow, dispatcher, event)

	output, err := useCase.Execute(ctx, input)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	mockUow.AssertExpectations(t)
	mockUow.AssertNumberOfCalls(t, "Do", 1)
}
