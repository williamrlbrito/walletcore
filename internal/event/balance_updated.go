package event

import "time"

type BalanceUpdatedEvent struct {
	Name    string
	Payload interface{}
}

func NewBalanceUpdatedEvent() *BalanceUpdatedEvent {
	return &BalanceUpdatedEvent{
		Name: "balance.updated",
	}
}

func (event *BalanceUpdatedEvent) GetName() string {
	return event.Name
}

func (event *BalanceUpdatedEvent) GetPayload() interface{} {
	return event.Payload
}

func (event *BalanceUpdatedEvent) GetDateTime() time.Time {
	return time.Now()
}

func (event *BalanceUpdatedEvent) SetPayload(payload interface{}) {
	event.Payload = payload
}
