package database

import (
	"database/sql"

	"github.com/emmanuelviniciusdev/imersao-fullcycle-microservicos/walletcore-api/internal/entity"
)

type TransactionDB struct {
	DB *sql.DB
}

func NewTransactionDB(db *sql.DB) *TransactionDB {
	return &TransactionDB{DB: db}
}

func (db *TransactionDB) Save(transaction *entity.Transaction) error {
	stmt, err := db.DB.Prepare(`
		INSERT INTO tb_transaction (id, account_id_from, account_id_to, amount, created_at)
		VALUES (?, ?, ?, ?, ?)
	`)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		transaction.ID,
		transaction.AccountFrom.ID,
		transaction.AccountTo.ID,
		transaction.Amount,
		transaction.CreatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}
