package entity_test

import (
	"testing"

	"github.com/emmanuelviniciusdev/imersao-fullcycle-microservicos/walletcore-api/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewTransactionInvalidAmount(t *testing.T) {
	customerRomeo, _ := entity.NewCustomer("Romeo", "romeo@icloud.com")
	customerDarwin, _ := entity.NewCustomer("Darwin", "darwin@icloud.com")

	accountRomeo := entity.NewAccount(customerRomeo, 150.5)
	accountDarwin := entity.NewAccount(customerDarwin, 50.5)

	transaction, err := entity.NewTransaction(accountRomeo, accountDarwin, -50.0)

	assert.EqualError(t, err, "amount must be greater than zero")

	assert.Nil(t, transaction)
}

func TestCreateNewTransactionInsufficientFunds(t *testing.T) {
	customerRomeo, _ := entity.NewCustomer("Romeo", "romeo@icloud.com")
	customerDarwin, _ := entity.NewCustomer("Darwin", "darwin@icloud.com")

	accountRomeo := entity.NewAccount(customerRomeo, 49.99)
	accountDarwin := entity.NewAccount(customerDarwin, 50.5)

	transaction, err := entity.NewTransaction(accountRomeo, accountDarwin, 50.0)

	assert.EqualError(t, err, "insufficient funds")

	assert.Nil(t, transaction)
}

func TestCreateNewTransaction(t *testing.T) {
	customerRomeo, _ := entity.NewCustomer("Romeo", "romeo@icloud.com")
	customerDarwin, _ := entity.NewCustomer("Darwin", "darwin@icloud.com")

	accountRomeo := entity.NewAccount(customerRomeo, 150.5)
	accountDarwin := entity.NewAccount(customerDarwin, 50.5)

	transaction, err := entity.NewTransaction(accountRomeo, accountDarwin, 50.0)

	assert.Nil(t, err)

	assert.Equal(t, 50.0, transaction.Amount)

	assert.Equal(t, 100.5, accountRomeo.Balance)
	assert.Equal(t, 100.5, accountDarwin.Balance)
}
