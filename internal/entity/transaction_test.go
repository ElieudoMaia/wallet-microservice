package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
	client1, _ := NewClient("John Doe", "j@j")
	account1, _ := NewAccount(client1)
	client2, _ := NewClient("Jane Doe 2", "j2@j")
	account2, _ := NewAccount(client2)

	account1.Credit(1000)
	account2.Credit(1000)

	transaction, err := NewTransaction(account1, account2, 100)
	assert.Nil(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, account1.ID, transaction.AccountFrom.ID)
	assert.Equal(t, account2.ID, transaction.AccountTo.ID)
	assert.Equal(t, 100.0, transaction.Amount)
	assert.Equal(t, 900.0, account1.Balance)
	assert.Equal(t, 1100.0, account2.Balance)
}

func TestCreateTransactionWithInsuficientBalance(t *testing.T) {
	client1, _ := NewClient("John Doe", "j@j")
	account1, _ := NewAccount(client1)
	client2, _ := NewClient("Jane Doe 2", "j2@j")
	account2, _ := NewAccount(client2)

	account1.Credit(1000)
	account2.Credit(1000)

	transaction, err := NewTransaction(account1, account2, 1001)
	assert.NotNil(t, err)
	assert.Nil(t, transaction)
	assert.Equal(t, "insufficient funds", err.Error())
	assert.Equal(t, 1000.0, account1.Balance)
	assert.Equal(t, 1000.0, account2.Balance)
}

func TestCreateTransactionWithNegativeAmount(t *testing.T) {
	client1, _ := NewClient("John Doe", "j@j")
	account1, _ := NewAccount(client1)
	client2, _ := NewClient("Jane Doe 2", "j2@j")
	account2, _ := NewAccount(client2)

	account1.Credit(1000)
	account2.Credit(1000)

	transaction, err := NewTransaction(account1, account2, -100)
	assert.NotNil(t, err)
	assert.Nil(t, transaction)
	assert.Equal(t, "amount must be greater than 0", err.Error())
	assert.Equal(t, 1000.0, account1.Balance)
	assert.Equal(t, 1000.0, account2.Balance)
}
