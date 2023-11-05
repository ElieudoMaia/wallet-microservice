package main

import (
	"context"
	"database/sql"
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/elieudomaia/ms-wallet-app/internal/database"
	"github.com/elieudomaia/ms-wallet-app/internal/event"
	handler "github.com/elieudomaia/ms-wallet-app/internal/event/handlers"
	"github.com/elieudomaia/ms-wallet-app/internal/usecase/create_account"
	"github.com/elieudomaia/ms-wallet-app/internal/usecase/create_client"
	"github.com/elieudomaia/ms-wallet-app/internal/usecase/create_transaction"
	"github.com/elieudomaia/ms-wallet-app/internal/web"
	"github.com/elieudomaia/ms-wallet-app/internal/web/webserver"
	"github.com/elieudomaia/ms-wallet-app/pkg/events"
	"github.com/elieudomaia/ms-wallet-app/pkg/kafka"
	"github.com/elieudomaia/ms-wallet-app/pkg/uow"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/wallet?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	configMap := ckafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "wallet",
	}
	kafkaProducer := kafka.NewKafkaProducer(&configMap)

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("TransactionCreated", handler.NewTransactionCreatedKafkaHandler(kafkaProducer))
	eventDispatcher.Register("BalanceUpdated", handler.NewUpdateBalanceKafkaHandler(kafkaProducer))
	transactionCreatedEvent := event.NewTransactionCreatedEvent()
	balanceUpdatedEvent := event.NewBalanceUpdatedEvent()

	clientDb := database.NewClientDB(db)
	accountDb := database.NewAccountDB(db)

	ctx := context.Background()
	uow := uow.NewUow(ctx, db)

	uow.Register("AccountDB", func(tx *sql.Tx) interface{} {
		return database.NewAccountDB(db)
	})

	uow.Register("TransactionDB", func(tx *sql.Tx) interface{} {
		return database.NewTransactionDB(db)
	})

	createTransactionUseCase := create_transaction.NewCreateTransactionUseCase(
		uow,
		eventDispatcher,
		transactionCreatedEvent,
		balanceUpdatedEvent,
	)
	createClientUseCase := create_client.NewCreateClientUseCase(clientDb)
	createAccountUseCase := create_account.NewCreateAccountUseCase(accountDb, clientDb)

	port := "8080"

	webserver := webserver.NewWebServer(fmt.Sprintf(":%s", port))

	clientHandler := web.NewWebClientHandler(*createClientUseCase)
	accountHandler := web.NewAccountHandler(*createAccountUseCase)
	transactionHandler := web.NewTransactionHandler(*createTransactionUseCase)

	webserver.AddHandler("/clients", clientHandler.CreateClient)
	webserver.AddHandler("/accounts", accountHandler.CreateAccount)
	webserver.AddHandler("/transactions", transactionHandler.CreateTransaction)

	fmt.Println("Server running on port", port)
	webserver.Start()
}
