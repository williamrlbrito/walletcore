package main

import (
	"context"
	"database/sql"
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
	"github.com/williamrlbrito/walletcore/internal/database"
	"github.com/williamrlbrito/walletcore/internal/event"
	"github.com/williamrlbrito/walletcore/internal/event/handler"
	"github.com/williamrlbrito/walletcore/internal/usecase/create_account"
	"github.com/williamrlbrito/walletcore/internal/usecase/create_client"
	"github.com/williamrlbrito/walletcore/internal/usecase/create_transaction"
	"github.com/williamrlbrito/walletcore/internal/web"
	"github.com/williamrlbrito/walletcore/internal/web/webserver"
	"github.com/williamrlbrito/walletcore/pkg/events"
	"github.com/williamrlbrito/walletcore/pkg/kafka"
	"github.com/williamrlbrito/walletcore/pkg/uow"
)

func main() {
	db, err := sql.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?parseTime=true",
			"root",
			"root",
			"mysql",
			"3306",
			"wallet",
		))

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
	eventDispatcher.Register("transaction.created", handler.NewTransactionCreatedKafkaHandler(kafkaProducer))
	eventDispatcher.Register("balance.updated", handler.NewBalanceUpdatedKafkaHandler(kafkaProducer))
	transactionCreated := event.NewTransactionCreatedEvent()
	balanceUpdated := event.NewBalanceUpdatedEvent()

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

	createClientUseCase := create_client.NewCreateClientUseCase(clientDb)
	createAccountUseCase := create_account.NewCreateAccountUseCase(accountDb, clientDb)
	createTransactionUseCase := create_transaction.NewCreateTransactionUseCase(
		uow,
		eventDispatcher,
		transactionCreated,
		balanceUpdated,
	)

	webServer := webserver.NewWebServer(":8080")
	clientHandler := web.NewWebClientHandler(*createClientUseCase)
	accountHandler := web.NewWebAccountHandler(*createAccountUseCase)
	transactionHandler := web.NewWebTransactionHandler(*createTransactionUseCase)

	webServer.AddHandler("/clients", clientHandler.CreateClient)
	webServer.AddHandler("/accounts", accountHandler.CreateAccount)
	webServer.AddHandler("/transactions", transactionHandler.CreateTransaction)

	fmt.Println("Server running on port 8080")
	webServer.Start()
}
