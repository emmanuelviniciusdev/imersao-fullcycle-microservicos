package usecase_test

import (
	"testing"

	"github.com/emmanuelviniciusdev/imersao-fullcycle-microservicos/walletcore-api/internal/entity"
	"github.com/emmanuelviniciusdev/imersao-fullcycle-microservicos/walletcore-api/internal/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type AccountGatewayMock struct {
	mock.Mock
}

func (mock *AccountGatewayMock) Save(account *entity.Account) error {
	args := mock.Called(account)
	return args.Error(0)
}

func (mock *AccountGatewayMock) Get(ID string) (*entity.Account, error) {
	args := mock.Called(ID)
	return args.Get(0).(*entity.Account), args.Error(1)
}

func TestCreateAccountUsecase_Execute(t *testing.T) {
	customer, _ := entity.NewCustomer("Emmanuel", "emmanuel@icloud.com")

	customerGatewayMock := &CustomerGatewayMock{}
	customerGatewayMock.On("Get", customer.ID).Return(customer, nil)

	accountGatewayMock := &AccountGatewayMock{}
	accountGatewayMock.On("Save", mock.Anything).Return(nil)

	uc := usecase.NewCreateAccountUsecase(accountGatewayMock, customerGatewayMock)

	input := usecase.CreateAccountInputDTO{CustomerID: customer.ID}

	output, err := uc.Execute(input)

	assert.Nil(t, err)

	assert.NotEmpty(t, output.ID)

	customerGatewayMock.AssertExpectations(t)
	customerGatewayMock.AssertNumberOfCalls(t, "Get", 1)

	accountGatewayMock.AssertExpectations(t)
	accountGatewayMock.AssertNumberOfCalls(t, "Save", 1)
}
