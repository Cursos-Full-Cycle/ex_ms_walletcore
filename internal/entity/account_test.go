package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "a@a.com")
	account := NewAccount(client)
	assert.NotNil(t, account)
	assert.Equal(t, client.ID, account.Client.ID)
	assert.Equal(t, 0.0, account.Balance)
}

func TestCreateNewAccountWhenClientIsNil(t *testing.T) {
	account := NewAccount(nil)
	assert.Nil(t, account)
}

func TestCreditAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "a@a.com")
	account := NewAccount(client)
	account.Credit(100)
	assert.Equal(t, float64(100), account.Balance)
}

func TestDebitAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "a@a.com")
	account := NewAccount(client)
	account.Credit(100)
	account.Debit(50)
	assert.Equal(t, float64(50), account.Balance)
}
