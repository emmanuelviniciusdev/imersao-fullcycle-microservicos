package usecase

import (
	"time"

	"github.com/emmanuelviniciusdev/imersao-fullcycle-microservicos/walletcore-api/internal/entity"
	"github.com/emmanuelviniciusdev/imersao-fullcycle-microservicos/walletcore-api/internal/gateway"
)

type CreateCustomerInputDTO struct {
	Name  string
	Email string
}

type CreateCustomerOutputDTO struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateCustomerUsecase struct {
	CustomerGateway gateway.CustomerGateway
}

func NewCreateCustomerUsecase(customerGateway gateway.CustomerGateway) *CreateCustomerUsecase {
	return &CreateCustomerUsecase{
		CustomerGateway: customerGateway,
	}
}

func (usecase *CreateCustomerUsecase) Execute(input CreateCustomerInputDTO) (*CreateCustomerOutputDTO, error) {
	entity, err := entity.NewCustomer(input.Name, input.Email)

	if err != nil {
		return nil, err
	}

	err = usecase.CustomerGateway.Save(entity)

	if err != nil {
		return nil, err
	}

	return &CreateCustomerOutputDTO{
		ID:        entity.ID,
		Name:      entity.Name,
		Email:     entity.Name,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}, nil
}
