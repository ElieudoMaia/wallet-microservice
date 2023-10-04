package main

import (
	"database/sql"
	"fmt"

	"github.com/elieudomaia/ms-wallet-app/internal/database"
	"github.com/elieudomaia/ms-wallet-app/internal/event"
	"github.com/elieudomaia/ms-wallet-app/internal/usecase/create_account"
	"github.com/elieudomaia/ms-wallet-app/internal/usecase/create_client"
	"github.com/elieudomaia/ms-wallet-app/internal/usecase/create_transaction"
	"github.com/elieudomaia/ms-wallet-app/internal/web"
	"github.com/elieudomaia/ms-wallet-app/internal/web/webserver"
	"github.com/elieudomaia/ms-wallet-app/pkg/events"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "mysql", "3306", "wallet"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	eventDispatcher := events.NewEventDispatcher()
	transactionCreatedEvent := event.NewTransactionCreatedEvent()
	// eventDispatcher.Register("TransactionCreated", handler)

	clientDb := database.NewClientDB(db)
	accountDb := database.NewAccountDB(db)
	transactionDb := database.NewTransactionDB(db)

	createClientUseCase := create_client.NewCreateClientUseCase(clientDb)
	createAccountUseCase := create_account.NewCreateAccountUseCase(accountDb, clientDb)
	createTransactionUseCase := create_transaction.NewCreateTransactionUseCase(transactionDb, accountDb, *eventDispatcher, transactionCreatedEvent)

	webserver := webserver.NewWebServer(":8080")

	clientHandler := web.NewWebClientHandler(*createClientUseCase)
	accountHandler := web.NewAccountHandler(*createAccountUseCase)
	transactionHandler := web.NewTransactionHandler(*createTransactionUseCase)

	webserver.AddHandler("/client", clientHandler.CreateClient)
	webserver.AddHandler("/account", accountHandler.CreateAccount)
	webserver.AddHandler("/transaction", transactionHandler.CreateTransaction)

	webserver.Start()
}
