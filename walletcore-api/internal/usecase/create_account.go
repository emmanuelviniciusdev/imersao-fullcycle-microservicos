package usecase

import (
	"github.com/emmanuelviniciusdev/imersao-fullcycle-microservicos/walletcore-api/internal/entity"
	"github.com/emmanuelviniciusdev/imersao-fullcycle-microservicos/walletcore-api/internal/gateway"
)

type CreateAccountInputDTO struct {
	CustomerID string
}

type CreateAccountOutputDTO struct {
	ID string
}

type CreateAccountUsecase struct {
	AccountGateway  gateway.AccountGateway
	CustomerGateway gateway.CustomerGateway
}

func NewCreateAccountUsecase(
	accountGateway gateway.AccountGateway,
	customerGateway gateway.CustomerGateway,
) *CreateAccountUsecase {
	return &CreateAccountUsecase{
		AccountGateway:  accountGateway,
		CustomerGateway: customerGateway,
	}
}

func (usecase *CreateAccountUsecase) Execute(input CreateAccountInputDTO) (*CreateAccountOutputDTO, error) {
	customer, err := usecase.CustomerGateway.Get(input.CustomerID)

	if err != nil {
		return nil, err
	}

	account := entity.NewAccount(customer, 0)

	err = usecase.AccountGateway.Save(account)

	if err != nil {
		return nil, err
	}

	return &CreateAccountOutputDTO{ID: account.ID}, nil
}
