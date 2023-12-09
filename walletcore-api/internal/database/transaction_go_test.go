package database_test

import (
	"database/sql"
	"testing"

	"github.com/emmanuelviniciusdev/imersao-fullcycle-microservicos/walletcore-api/internal/database"
	"github.com/emmanuelviniciusdev/imersao-fullcycle-microservicos/walletcore-api/internal/entity"
	"github.com/stretchr/testify/suite"
)

type TransactionDBTestSuite struct {
	suite.Suite

	DB            *sql.DB
	TransactionDB *database.TransactionDB
}

func (testSuite *TransactionDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")

	testSuite.Nil(err)

	db.Exec("CREATE TABLE IF NOT EXISTS tb_transaction (id varchar(255), account_id_from varchar(255), account_id_to varchar(255), amount float, created_at datetime)")

	testSuite.DB = db
	testSuite.TransactionDB = database.NewTransactionDB(db)
}

func (testSuite *TransactionDBTestSuite) TearDownSuite() {
	testSuite.DB.Exec("DROP TABLE tb_transaction")
	testSuite.DB.Close()
}

func TestTransactionDBTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionDBTestSuite))
}

func (testSuite *TransactionDBTestSuite) TestSave() {
	customerRomeo, _ := entity.NewCustomer("Romeo", "romeo@icloud.com")
	accountRomeo := entity.NewAccount(customerRomeo, 150)

	customerDarwin, _ := entity.NewCustomer("Darwin", "darwin@icloud.com")
	accountDarwin := entity.NewAccount(customerDarwin, 0)

	transaction, _ := entity.NewTransaction(accountRomeo, accountDarwin, 50.0)

	err := testSuite.TransactionDB.Save(transaction)

	testSuite.Assert().Nil(err)
}
