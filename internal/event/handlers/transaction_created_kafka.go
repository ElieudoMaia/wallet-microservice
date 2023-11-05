package handler

import (
	"fmt"
	"sync"

	"github.com/elieudomaia/ms-wallet-app/pkg/events"
	"github.com/elieudomaia/ms-wallet-app/pkg/kafka"
)

type TransactionCreatedKafkaHandler struct {
	Kafka *kafka.Producer
}

func NewTransactionCreatedKafkaHandler(kafka *kafka.Producer) *TransactionCreatedKafkaHandler {
	return &TransactionCreatedKafkaHandler{
		Kafka: kafka,
	}
}

func (h *TransactionCreatedKafkaHandler) Handle(message events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	err := h.Kafka.Publish(message, nil, "transactions")
	if err != nil {
		panic(err)
	}
	fmt.Println("TransactionCreatedKafkaHandler: ", message.GetPayload())
}
