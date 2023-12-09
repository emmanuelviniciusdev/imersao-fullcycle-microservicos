package database_test

import (
	"database/sql"
	"testing"

	"github.com/emmanuelviniciusdev/imersao-fullcycle-microservicos/walletcore-api/internal/database"
	"github.com/emmanuelviniciusdev/imersao-fullcycle-microservicos/walletcore-api/internal/entity"
	"github.com/stretchr/testify/suite"

	_ "github.com/mattn/go-sqlite3"
)

type CustomerDBTestSuite struct {
	suite.Suite

	DB         *sql.DB
	CustomerDB *database.CustomerDB
}

func (testSuite *CustomerDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")

	testSuite.Nil(err)

	db.Exec("CREATE TABLE IF NOT EXISTS tb_customer (id varchar(255), name varchar(255), email varchar(255), created_at datetime, updated_at datetime)")

	testSuite.DB = db
	testSuite.CustomerDB = database.NewCustomerDB(db)
}

func (testSuite *CustomerDBTestSuite) TearDownSuite() {
	testSuite.DB.Exec("DROP TABLE tb_customer")
	testSuite.DB.Close()
}

func TestCustomerDBTestSuite(t *testing.T) {
	suite.Run(t, new(CustomerDBTestSuite))
}

func (testSuite *CustomerDBTestSuite) TestSave() {
	customer, _ := entity.NewCustomer("Emmanuel", "emmanuel@icloud.com")

	persistedCustomer, err := testSuite.CustomerDB.Save(customer)

	testSuite.Assert().Nil(err)
	testSuite.Assert().NotEmpty(persistedCustomer)
}

func (testSuite *CustomerDBTestSuite) TestGet() {
	customer, _ := entity.NewCustomer("Emmanuel", "emmanuel@icloud.com")

	testSuite.CustomerDB.Save(customer)

	foundCustomer, err := testSuite.CustomerDB.Get(customer.ID)

	testSuite.Assert().Nil(err)

	testSuite.Assert().Equal(customer.ID, foundCustomer.ID)
	testSuite.Assert().Equal(customer.Name, foundCustomer.Name)
	testSuite.Assert().Equal(customer.Email, foundCustomer.Email)

	testSuite.Assert().NotEmpty(foundCustomer.CreatedAt)
	testSuite.Assert().NotEmpty(foundCustomer.UpdatedAt)
}
