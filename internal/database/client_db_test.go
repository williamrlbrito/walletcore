package database

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
	"github.com/williamrlbrito/walletcore/internal/entity"
)

type ClientDBTestSuite struct {
	suite.Suite
	db       *sql.DB
	ClientDB *ClientDB
}

func (suite *ClientDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.Nil(err)
	suite.db = db
	db.Exec("CREATE TABLE clients (id TEXT, name TEXT, email TEXT, created_at DATETIME)")
	suite.ClientDB = NewClientDB(db)
}

func (suite *ClientDBTestSuite) TearDownSuite() {
	defer suite.db.Close()
	suite.db.Exec("DROP TABLE clients")
}

func TestClientDBTestSuite(t *testing.T) {
	suite.Run(t, new(ClientDBTestSuite))
}

func (suite *ClientDBTestSuite) TestSave() {
	client, _ := entity.NewClient("John Doe", "john@doe.com")
	err := suite.ClientDB.Save(client)
	suite.Nil(err)
}

func (suite *ClientDBTestSuite) TestGet() {
	client, _ := entity.NewClient("John Doe", "john@doe.com")
	suite.ClientDB.Save(client)

	clientDB, err := suite.ClientDB.Get(client.ID)
	suite.Nil(err)
	suite.Equal(client.ID, clientDB.ID)
	suite.Equal(client.Name, clientDB.Name)
	suite.Equal(client.Email, clientDB.Email)
}
