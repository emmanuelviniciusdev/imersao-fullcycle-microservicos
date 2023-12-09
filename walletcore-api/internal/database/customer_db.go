package database

import (
	"database/sql"

	"github.com/emmanuelviniciusdev/imersao-fullcycle-microservicos/walletcore-api/internal/entity"
)

type CustomerDB struct {
	DB *sql.DB
}

func NewCustomerDB(db *sql.DB) *CustomerDB {
	return &CustomerDB{DB: db}
}

func (db *CustomerDB) Save(customer *entity.Customer) (*entity.Customer, error) {
	stmt, err := db.DB.Prepare("INSERT INTO tb_customer (id, name, email, created_at, updated_at) VALUES (?, ?, ?, ?, ?)")

	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(customer.ID, customer.Name, customer.Email, customer.CreatedAt, customer.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (db *CustomerDB) Get(ID string) (*entity.Customer, error) {
	customer := &entity.Customer{}

	stmt, err := db.DB.Prepare("SELECT id, name, email, created_at, updated_at FROM tb_customer WHERE id = ?")

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(ID)

	if err = row.Scan(&customer.ID, &customer.Name, &customer.Email, &customer.CreatedAt, &customer.UpdatedAt); err != nil {
		return nil, err
	}

	return customer, nil
}
