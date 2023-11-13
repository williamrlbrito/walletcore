package walletcore

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/williamrlbrito/walletcore/internal/database"
	"github.com/williamrlbrito/walletcore/internal/event"
	createaccount "github.com/williamrlbrito/walletcore/internal/usecase/create_account"
	createclient "github.com/williamrlbrito/walletcore/internal/usecase/create_client"
	createtransaction "github.com/williamrlbrito/walletcore/internal/usecase/create_transaction"
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

	createClientUseCase := createclient.NewCreateClientUseCase(clientDb)
	createAccountUseCase := createaccount.NewCreateAccountUseCase(accountDb, clientDb)
	createTransactionUseCase := createtransaction.NewCreateTransactionUseCase(
		transactionDb,
		accountDb,
		eventDispatcher,
		transactionCreated)
}
