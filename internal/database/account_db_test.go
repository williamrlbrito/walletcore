package database

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/williamrlbrito/walletcore/internal/entity"
)

type AccountDBTestSuite struct {
	suite.Suite
	db        *sql.DB
	accountDB *AccountDB
	client    *entity.Client
}

func (suite *AccountDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.Nil(err)
	suite.db = db
	db.Exec("CREATE TABLE clients (id TEXT, name TEXT, email TEXT, created_at DATETIME)")
	db.Exec("CREATE TABLE accounts (id TEXT, client_id TEXT, balance float, created_at DATETIME)")
	suite.accountDB = NewAccountDB(db)
	suite.client, _ = entity.NewClient("John Doe", "john@doe.com")
}

func (suite *AccountDBTestSuite) TearDownSuite() {
	defer suite.db.Close()
	suite.db.Exec("DROP TABLE clients")
	suite.db.Exec("DROP TABLE accounts")
}

func TestAccountDBTestSuite(t *testing.T) {
	suite.Run(t, new(AccountDBTestSuite))
}
