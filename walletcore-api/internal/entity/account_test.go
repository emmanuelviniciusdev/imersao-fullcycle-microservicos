package entity_test

import (
	"testing"

	"github.com/emmanuelviniciusdev/imersao-fullcycle-microservicos/walletcore-api/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestDebitAccount(t *testing.T) {
	customer, _ := entity.NewCustomer("Romeo", "romeo@icloud.com")

	account := entity.NewAccount(customer, 100.5)

	account.Debit(0.5)

	assert.Equal(t, 100.0, account.Balance)
}

func TestCreditAccount(t *testing.T) {
	customer, _ := entity.NewCustomer("Romeo", "romeo@icloud.com")

	account := entity.NewAccount(customer, 100.5)

	account.Credit(0.5)

	assert.Equal(t, 101.0, account.Balance)
}

func TestCreateNewAccountInvalidCustomer(t *testing.T) {
	account := entity.NewAccount(nil, 1_350_700.97)
	assert.Nil(t, account)
}

func TestCreateNewAccount(t *testing.T) {
	customer, _ := entity.NewCustomer("Romeo", "romeo@icloud.com")

	account := entity.NewAccount(customer, 1_350_700.97)

	assert.NotNil(t, account)

	assert.Equal(t, customer, account.Customer)
	assert.Equal(t, 1_350_700.97, account.Balance)

	assert.NotNil(t, account.ID)
	assert.NotNil(t, account.CreatedAt)
	assert.NotNil(t, account.UpdatedAt)
}
