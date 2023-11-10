package database

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/williamrlbrito/walletcore/internal/entity"
)

type TransactionDBTestSuite struct {
	suite.Suite
	db            *sql.DB
	clientFrom    *entity.Client
	clientTo      *entity.Client
	accountFrom   *entity.Account
	accountTo     *entity.Account
	transactionDB *TransactionDB
}

func (suite *TransactionDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.Nil(err)
	suite.db = db
	db.Exec("CREATE TABLE clients (id TEXT, name TEXT, email TEXT, created_at DATETIME)")
	db.Exec("CREATE TABLE accounts (id TEXT, client_id TEXT, balance float, created_at DATETIME)")
	db.Exec("CREATE TABLE transactions (id TEXT, account_id_from TEXT, account_id_to TEXT, amount float, created_at DATETIME)")
	clientFrom, err := entity.NewClient("John Doe", "john@doe.com")
	suite.Nil(err)
	suite.clientFrom = clientFrom

	clientTo, err := entity.NewClient("Jane Doe", "jane@doe.com")
	suite.Nil(err)
	suite.clientTo = clientTo

	accountFrom := entity.NewAccount(clientFrom)
	accountFrom.Credit(1000)
	suite.accountFrom = accountFrom

	accountTo := entity.NewAccount(clientTo)
	accountTo.Credit(1000)
	suite.accountTo = accountTo

	suite.transactionDB = NewTransactionDB(db)
}

func (suite *TransactionDBTestSuite) TearDownSuite() {
	defer suite.db.Close()
	suite.db.Exec("DROP TABLE clients")
	suite.db.Exec("DROP TABLE accounts")
	suite.db.Exec("DROP TABLE transactions")
}

func TestTransactionDBTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionDBTestSuite))
}

func (suite *TransactionDBTestSuite) TestCreate() {
	transaction, err := entity.NewTransaction(suite.accountFrom, suite.accountTo, 100)
	suite.Nil(err)

	err = suite.transactionDB.Create(transaction)
	suite.Nil(err)
}
