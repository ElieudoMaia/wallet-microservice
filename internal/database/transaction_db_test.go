package database

import (
	"database/sql"
	"testing"

	"github.com/elieudomaia/ms-wallet-app/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TransactionDBTestSuite struct {
	suite.Suite
	db            *sql.DB
	client1       *entity.Client
	client2       *entity.Client
	accountFrom   *entity.Account
	accountTo     *entity.Account
	transactionDB *TransactionDB
}

func (s *TransactionDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	assert.Nil(s.T(), err)
	s.db = db
	db.Exec("CREATE TABLE clients (id varchar(255), name varchar(255), email varchar(255), created_at DATETIME)")
	db.Exec("CREATE TABLE accounts (id varchar(255), client_id varchar(255), balance int, created_at DATETIME)")
	db.Exec("CREATE TABLE transactions (id varchar(255), account_id_from varchar(255), account_id_to varchar(255), amount real, created_at DATETIME)")

	client1, err := entity.NewClient("John Doe", "j@j.com")
	assert.Nil(s.T(), err)
	s.client1 = client1
	client2, err := entity.NewClient("Jane Doe", "j2@j.com")
	assert.Nil(s.T(), err)
	s.client2 = client2
	accountFrom, err := entity.NewAccount(s.client1)
	assert.Nil(s.T(), err)
	accountFrom.Balance = 1000.0
	s.accountFrom = accountFrom
	accountTo, err := entity.NewAccount(s.client2)
	assert.Nil(s.T(), err)
	accountTo.Balance = 1000.0
	s.accountTo = accountTo

	s.transactionDB = NewTransactionDB(db)
}

func (s *TransactionDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
	s.db.Exec("DROP TABLE accounts")
	s.db.Exec("DROP TABLE transactions")
}

func TestTransactionDBTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionDBTestSuite))
}

func (s *TransactionDBTestSuite) TestCreate() {
	transaction, err := entity.NewTransaction(s.accountFrom, s.accountTo, 100.0)
	assert.Nil(s.T(), err)
	err = s.transactionDB.Create(transaction)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), s.accountFrom.Balance, 900.0)
	assert.Equal(s.T(), s.accountTo.Balance, 1100.0)
}
