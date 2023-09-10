package event

import "time"

type TransactionCreatedEvent struct {
	Name     string
	DateTime time.Time
	Payload  interface{}
}

func NewTransactionCreatedEvent() *TransactionCreatedEvent {
	return &TransactionCreatedEvent{
		Name:     "TransactionCreated",
		DateTime: time.Now(),
	}
}

func (e *TransactionCreatedEvent) GetName() string {
	return e.Name
}

func (e *TransactionCreatedEvent) GetDateTime() time.Time {
	return e.DateTime
}

func (e *TransactionCreatedEvent) GetPayload() interface{} {
	return e.Payload
}

func (e *TransactionCreatedEvent) SetPayload(payload interface{}) {
	e.Payload = payload
}
