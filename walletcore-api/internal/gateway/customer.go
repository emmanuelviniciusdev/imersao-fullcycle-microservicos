package gateway

import "github.com/emmanuelviniciusdev/imersao-fullcycle-microservicos/walletcore-api/internal/entity"

type CustomerGateway interface {
	Get(ID string) (*entity.Customer, error)
	Save(customer *entity.Customer) error
}
