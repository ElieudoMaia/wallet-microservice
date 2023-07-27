package database

import (
	"database/sql"
	"testing"

	"github.com/elieudomaia/ms-wallet-app/internal/entity"
	"github.com/stretchr/testify/suite"
)

type AccountDBTestSuite struct {
	suite.Suite
	db        *sql.DB
	accountDB *AccountDB
	client    *entity.Client
}

func (s *AccountDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.NotNil(db)
	s.db = db
	db.Exec("CREATE TABLE clients (id varchar(255), name varchar(255), email varchar(255), created_at DATETIME)")
	db.Exec("CREATE TABLE accounts (id varchar(255), client_id varchar(255), balance int, created_at DATETIME)")
	s.accountDB = NewAccountDB(db)
	s.client, _ = entity.NewClient("John Doe", "mail@mail.com")

}

func (s *AccountDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
	s.db.Exec("DROP TABLE accounts")
}

func TestAccountDBTestSuite(t *testing.T) {
	suite.Run(t, new(AccountDBTestSuite))
}

func (s *AccountDBTestSuite) TestSave() {
	account, _ := entity.NewAccount(s.client)
	err := s.accountDB.Save(account)
	s.Nil(err)
}

func (s *AccountDBTestSuite) TestFindByID() {
	s.db.Exec(
		`insert into clients (id, name, email, created_at) values (?,?,?,?)`,
		s.client.ID, s.client.Name, s.client.Email, s.client.CreatedAt,
	)
	account, _ := entity.NewAccount(s.client)
	err := s.accountDB.Save(account)
	s.Nil(err)
	acc, er := s.accountDB.FindByID(account.ID)
	s.Nil(er)
	s.Equal(account.ID, acc.ID)
	s.Equal(account.Client.ID, acc.Client.ID)
	s.Equal(account.Balance, acc.Balance)
}
