package main

import (
	"context"
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
	"github.com/elieudomaia/ms-wallet-app/pkg/uow"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/wallet?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	eventDispatcher := events.NewEventDispatcher()
	transactionCreatedEvent := event.NewTransactionCreatedEvent()
	// eventDispatcher.Register("TransactionCreated", handler)

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

	createTransactionUseCase := create_transaction.NewCreateTransactionUseCase(uow, eventDispatcher, transactionCreatedEvent)
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
