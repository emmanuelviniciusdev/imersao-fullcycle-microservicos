package database

import (
	"database/sql"

	"github.com/emmanuelviniciusdev/imersao-fullcycle-microservicos/walletcore-api/internal/entity"
)

type AccountDB struct {
	DB *sql.DB
}

func NewAccountDB(db *sql.DB) *AccountDB {
	return &AccountDB{DB: db}
}

func (db *AccountDB) Save(account *entity.Account) error {
	stmt, err := db.DB.Prepare("INSERT INTO tb_account (id, customer_id, balance, created_at, updated_at) VALUES (?, ?, ?, ?, ?)")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(account.ID, account.Customer.ID, account.Balance, account.CreatedAt, account.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (db *AccountDB) Get(ID string) (*entity.Account, error) {
	var account entity.Account
	var customer entity.Customer

	account.Customer = &customer

	stmt, err := db.DB.Prepare(`
		SELECT 
			ta.id, ta.balance, ta.created_at, ta.updated_at, 
			tc.id, tc.name, tc.email, tc.created_at, tc.updated_at 
		FROM tb_account ta
		JOIN tb_customer tc ON tc.id = ta.customer_id
		WHERE ta.id = ?
	`)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(ID)

	err = row.Scan(
		&account.ID,
		&account.Balance,
		&account.CreatedAt,
		&account.UpdatedAt,
		&customer.ID,
		&customer.Name,
		&customer.Email,
		&customer.CreatedAt,
		&customer.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &account, nil
}
