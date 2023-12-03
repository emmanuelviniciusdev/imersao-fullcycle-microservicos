package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	ID        string
	Email     string
	Name      string
	Accounts  []*Account
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewCustomer(name string, email string) (*Customer, error) {
	customer := &Customer{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := customer.Validate()

	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (c *Customer) AddAccount(account *Account) error {
	if account.Customer.ID != c.ID {
		return errors.New("the account does not belong to the customer")
	}

	c.Accounts = append(c.Accounts, account)

	return nil
}

func (c *Customer) Update(name string, email string) error {
	c.Name = name
	c.Email = email
	c.UpdatedAt = time.Now()

	err := c.Validate()

	return err
}

func (c *Customer) Validate() error {
	if c.Name == "" {
		return errors.New("the customer must have a name")
	}

	if c.Email == "" {
		return errors.New("the customer must have an email")
	}

	return nil
}
