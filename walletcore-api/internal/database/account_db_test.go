package database_test

import (
	"database/sql"
	"testing"

	"github.com/emmanuelviniciusdev/imersao-fullcycle-microservicos/walletcore-api/internal/database"
	"github.com/emmanuelviniciusdev/imersao-fullcycle-microservicos/walletcore-api/internal/entity"
	"github.com/stretchr/testify/suite"
)

type AccountDBTestSuite struct {
	suite.Suite

	DB         *sql.DB
	AccountDB  *database.AccountDB
	CustomerDB *database.CustomerDB
}

func (testSuite *AccountDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")

	testSuite.Nil(err)

	db.Exec("CREATE TABLE IF NOT EXISTS tb_customer (id varchar(255), name varchar(255), email varchar(255), created_at datetime, updated_at datetime)")
	db.Exec("CREATE TABLE IF NOT EXISTS tb_account (id varchar(255), customer_id varchar(255), balance float, created_at datetime, updated_at datetime)")

	testSuite.DB = db
	testSuite.AccountDB = database.NewAccountDB(db)
	testSuite.CustomerDB = database.NewCustomerDB(db)
}

func (testSuite *AccountDBTestSuite) TearDownSuite() {
	testSuite.DB.Exec("DROP TABLE tb_customer")
	testSuite.DB.Exec("DROP TABLE tb_account")
	testSuite.DB.Close()
}

func TestAccountDBTestSuite(t *testing.T) {
	suite.Run(t, new(AccountDBTestSuite))
}

func (testSuite *AccountDBTestSuite) TestSave() {
	customer, _ := entity.NewCustomer("Emmanuel", "emmanuel@icloud.com")
	account := entity.NewAccount(customer, 0)

	err := testSuite.AccountDB.Save(account)

	testSuite.Assert().Nil(err)
}

func (testSuite *AccountDBTestSuite) TestGet() {
	customer, _ := entity.NewCustomer("Emmanuel", "emmanuel@icloud.com")
	account := entity.NewAccount(customer, 0)

	_, err := testSuite.CustomerDB.Save(customer)
	testSuite.Assert().Nil(err)

	err = testSuite.AccountDB.Save(account)
	testSuite.Assert().Nil(err)

	persistedAccount, err := testSuite.AccountDB.Get(account.ID)

	testSuite.Assert().Nil(err)
	testSuite.Assert().Equal(account.Balance, persistedAccount.Balance)
	testSuite.Assert().Equal(account.Customer.ID, persistedAccount.Customer.ID)
	testSuite.Assert().Equal(account.Customer.Name, persistedAccount.Customer.Name)
	testSuite.Assert().Equal(account.Customer.Email, persistedAccount.Customer.Email)
	testSuite.Assert().NotEmpty(persistedAccount.CreatedAt)
	testSuite.Assert().NotEmpty(persistedAccount.UpdatedAt)
}
