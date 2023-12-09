package entity_test

import (
	"testing"

	"github.com/emmanuelviniciusdev/imersao-fullcycle-microservicos/walletcore-api/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestAddAccount(t *testing.T) {
	customerRomeo, _ := entity.NewCustomer("Romeo", "romeo@icloud.com")
	accountRomeo := entity.NewAccount(customerRomeo, 0)

	err := customerRomeo.AddAccount(accountRomeo)

	assert.Nil(t, err)

	assert.Equal(t, accountRomeo, customerRomeo.Accounts[0])

	customerDarwin, _ := entity.NewCustomer("Darwin", "darwin@icloud.com")

	err = customerDarwin.AddAccount(accountRomeo)

	assert.EqualError(t, err, "the account does not belong to the customer")
}

func TestUpdateCustomer(t *testing.T) {
	customer, _ := entity.NewCustomer("Romeo", "romeo@icloud.com")

	err := customer.Update("Romeo Montague", "romeo.montague.1597@icloud.com")

	assert.Nil(t, err)

	assert.Equal(t, "Romeo Montague", customer.Name)
	assert.Equal(t, "romeo.montague.1597@icloud.com", customer.Email)

	assert.NotEqual(t, customer.CreatedAt, customer.UpdatedAt)
}

func TestCreateNewCustomerInvalidEmail(t *testing.T) {
	customer, err := entity.NewCustomer("Romeo", "")

	assert.EqualError(t, err, "the customer must have an email")

	assert.Nil(t, customer)
}

func TestCreateNewCustomerInvalidName(t *testing.T) {
	customer, err := entity.NewCustomer("", "romeo@icloud.com")

	assert.EqualError(t, err, "the customer must have a name")

	assert.Nil(t, customer)
}

func TestCreateNewCustomer(t *testing.T) {
	customer, err := entity.NewCustomer("Romeo", "romeo@icloud.com")

	assert.Nil(t, err)

	assert.Equal(t, "Romeo", customer.Name)
	assert.Equal(t, "romeo@icloud.com", customer.Email)

	assert.NotNil(t, customer.ID)
	assert.NotNil(t, customer.CreatedAt)
	assert.NotNil(t, customer.UpdatedAt)
}
