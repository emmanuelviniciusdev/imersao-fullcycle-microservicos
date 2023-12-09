package usecase_test

import (
	"testing"

	"github.com/emmanuelviniciusdev/imersao-fullcycle-microservicos/walletcore-api/internal/entity"
	"github.com/emmanuelviniciusdev/imersao-fullcycle-microservicos/walletcore-api/internal/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type TransactionGatewayMock struct {
	mock.Mock
}

func (mock *TransactionGatewayMock) Save(transaction *entity.Transaction) error {
	args := mock.Called(transaction)
	return args.Error(0)
}

func TestCreateTransactionUsecase_Execute(t *testing.T) {
	customerRomeo, _ := entity.NewCustomer("Romeo", "romeo@icloud.com")
	accountRomeo := entity.NewAccount(customerRomeo, 150)

	customerDarwin, _ := entity.NewCustomer("Darwin", "darwin@icloud.com")
	accountDarwin := entity.NewAccount(customerDarwin, 50)

	accountGatewayMock := &AccountGatewayMock{}
	accountGatewayMock.On("Get", accountRomeo.ID).Return(accountRomeo, nil)
	accountGatewayMock.On("Get", accountDarwin.ID).Return(accountDarwin, nil)

	transactionGatewayMock := &TransactionGatewayMock{}
	transactionGatewayMock.On("Save", mock.Anything).Return(nil)

	uc := usecase.NewCreateTransactionUsecase(transactionGatewayMock, accountGatewayMock)

	input := usecase.CreateTransactionInputDTO{
		AccountIDFrom: accountRomeo.ID,
		AccountIDTo:   accountDarwin.ID,
		Amount:        50,
	}

	output, err := uc.Execute(input)

	assert.Nil(t, err)

	assert.NotEmpty(t, output.ID)

	accountGatewayMock.AssertExpectations(t)
	accountGatewayMock.AssertNumberOfCalls(t, "Get", 2)

	transactionGatewayMock.AssertExpectations(t)
	transactionGatewayMock.AssertNumberOfCalls(t, "Save", 1)
}
