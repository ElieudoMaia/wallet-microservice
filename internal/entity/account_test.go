package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "any@mail.com")
	account, err := NewAccount(client)
	assert.NotNil(t, account)
	assert.Nil(t, err)
	assert.Equal(t, client.ID, account.Client.ID)
}

func TestNewAccountEmptyClient(t *testing.T) {
	account, err := NewAccount(nil)
	assert.Nil(t, account)
	assert.NotNil(t, err)
	assert.Equal(t, "client is required", err.Error())
}

func TestCreditAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "any@mail.com")
	account, _ := NewAccount(client)
	err := account.Credit(100)
	assert.Nil(t, err)
	assert.Equal(t, 100.0, account.Balance)
}

func TestCreditInvalidAmount(t *testing.T) {
	client, _ := NewClient("John Doe", "any@mail.com")
	account, _ := NewAccount(client)
	err := account.Credit(0.0)
	assert.NotNil(t, err)
	assert.Equal(t, "amount must be greater than 0", err.Error())
	assert.Equal(t, 0.0, account.Balance)
}

func TestDebitAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "any@mail.com")
	account, _ := NewAccount(client)
	account.Credit(100)
	err := account.Debit(50)
	assert.Nil(t, err)
	assert.Equal(t, 50.0, account.Balance)
}

func TestDebitInvalidAmount(t *testing.T) {
	client, _ := NewClient("John Doe", "any@mail.com")
	account, _ := NewAccount(client)
	account.Credit(100.0)
	err := account.Debit(101.0)
	assert.NotNil(t, err)
	assert.Equal(t, 100.0, account.Balance)
}
