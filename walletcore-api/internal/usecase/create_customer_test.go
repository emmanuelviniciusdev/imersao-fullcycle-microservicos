package usecase_test

import (
	"testing"

	"github.com/emmanuelviniciusdev/imersao-fullcycle-microservicos/walletcore-api/internal/entity"
	"github.com/emmanuelviniciusdev/imersao-fullcycle-microservicos/walletcore-api/internal/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type CustomerGatewayMock struct {
	mock.Mock
}

func (mock *CustomerGatewayMock) Save(customer *entity.Customer) error {
	args := mock.Called(customer)
	return args.Error(0)
}

func (mock *CustomerGatewayMock) Get(ID string) (*entity.Customer, error) {
	args := mock.Called(ID)
	return args.Get(0).(*entity.Customer), args.Error(1)
}

func TestCreateCustomerUsecase_Execute(t *testing.T) {
	customerGatewayMock := &CustomerGatewayMock{}

	customerGatewayMock.On("Save", mock.Anything).Return(nil)

	uc := usecase.NewCreateCustomerUsecase(customerGatewayMock)

	input := usecase.CreateCustomerInputDTO{Name: "Emmanuel", Email: "emmanuel@icloud.com"}

	output, err := uc.Execute(input)

	assert.Nil(t, err)

	assert.Equal(t, "Emmanuel", output.Name)
	assert.Equal(t, "emmanuel@icloud.com", output.Email)

	assert.NotEmpty(t, output.ID)
	assert.NotEmpty(t, output.CreatedAt)
	assert.NotEmpty(t, output.UpdatedAt)

	customerGatewayMock.AssertExpectations(t)
	customerGatewayMock.AssertNumberOfCalls(t, "Save", 1)
}
