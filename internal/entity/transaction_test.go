package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
	clientFrom, _ := NewClient("John Doe", "j@j")
	accountFrom := NewAccount(clientFrom)
	clientTo, _ := NewClient("Jane Doe", "j@j")
	accountTo := NewAccount(clientTo)

	accountFrom.Credit(1000)
	accountTo.Credit(1000)

	transaction, err := NewTransaction(accountFrom, accountTo, 100)
	assert.Nil(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, accountFrom.Balance, 900.0)
	assert.Equal(t, accountTo.Balance, 1100.0)
}

func TestCreateTransactionWithInsufficientFunds(t *testing.T) {
	clientFrom, _ := NewClient("John Doe", "j@j")
	accountFrom := NewAccount(clientFrom)
	clientTo, _ := NewClient("Jane Doe", "j@j")
	accountTo := NewAccount(clientTo)

	accountFrom.Credit(1000)
	accountTo.Credit(1000)

	transaction, err := NewTransaction(accountFrom, accountTo, 2000)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "account from has insufficient funds")
	assert.Nil(t, transaction)
	assert.Equal(t, accountFrom.Balance, 1000.0)
	assert.Equal(t, accountTo.Balance, 1000.0)
}
