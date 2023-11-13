package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/williamrlbrito/walletcore/internal/database"
	"github.com/williamrlbrito/walletcore/internal/event"
	"github.com/williamrlbrito/walletcore/internal/usecase/create_account"
	"github.com/williamrlbrito/walletcore/internal/usecase/create_client"
	"github.com/williamrlbrito/walletcore/internal/usecase/create_transaction"
	"github.com/williamrlbrito/walletcore/internal/web"
	"github.com/williamrlbrito/walletcore/internal/web/webserver"
	"github.com/williamrlbrito/walletcore/pkg/events"
)

func main() {
	db, err := sql.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?parseTime=true",
			"root",
			"root",
			"localhost",
			"3306",
			"wallet",
		))

	if err != nil {
		panic(err)
	}

	defer db.Close()

	eventDispatcher := events.NewEventDispatcher()
	transactionCreated := event.NewTransactionCreatedEvent()
	// eventDispatcher.Register("transaction.created", handle)

	clientDb := database.NewClientDB(db)
	accountDb := database.NewAccountDB(db)
	transactionDb := database.NewTransactionDB(db)

	createClientUseCase := create_client.NewCreateClientUseCase(clientDb)
	createAccountUseCase := create_account.NewCreateAccountUseCase(accountDb, clientDb)
	createTransactionUseCase := create_transaction.NewCreateTransactionUseCase(
		transactionDb,
		accountDb,
		eventDispatcher,
		transactionCreated)

	webServer := webserver.NewWebServer(":3000")
	clientHandler := web.NewWebClientHandler(*createClientUseCase)
	accountHandler := web.NewWebAccountHandler(*createAccountUseCase)
	transactionHandler := web.NewWebTransactionHandler(*createTransactionUseCase)

	webServer.AddHandler("/clients", clientHandler.CreateClient)
	webServer.AddHandler("/accounts", accountHandler.CreateAccount)
	webServer.AddHandler("/transactions", transactionHandler.CreateTransaction)

	webServer.Start()
}
