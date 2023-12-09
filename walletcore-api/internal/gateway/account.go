package gateway

import "github.com/emmanuelviniciusdev/imersao-fullcycle-microservicos/walletcore-api/internal/entity"

type AccountGateway interface {
	Get(ID string) (*entity.Account, error)
	Save(account *entity.Account) error
}
