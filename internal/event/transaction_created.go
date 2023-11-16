package event

import "time"

type TransactionCreatedEvent struct {
	Name    string
	Payload interface{}
}

func NewTransactionCreatedEvent() *TransactionCreatedEvent {
	return &TransactionCreatedEvent{
		Name: "transaction.created",
	}
}

func (event *TransactionCreatedEvent) GetName() string {
	return event.Name
}

func (event *TransactionCreatedEvent) GetPayload() interface{} {
	return event.Payload
}

func (event *TransactionCreatedEvent) GetDateTime() time.Time {
	return time.Now()
}

func (event *TransactionCreatedEvent) SetPayload(payload interface{}) {
	event.Payload = payload
}
