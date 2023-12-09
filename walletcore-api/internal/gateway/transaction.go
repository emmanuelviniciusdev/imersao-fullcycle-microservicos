package gateway

import "github.com/emmanuelviniciusdev/imersao-fullcycle-microservicos/walletcore-api/internal/entity"

type TransactionGateway interface {
	Save(transaction *entity.Transaction) error
}
