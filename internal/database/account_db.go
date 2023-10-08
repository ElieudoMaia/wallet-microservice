package database

import (
	"database/sql"

	"github.com/elieudomaia/ms-wallet-app/internal/entity"
)

type AccountDB struct {
	DB *sql.DB
}

func NewAccountDB(db *sql.DB) *AccountDB {
	return &AccountDB{DB: db}
}

func (a *AccountDB) FindByID(id string) (*entity.Account, error) {
	var account entity.Account
	var client entity.Client
	account.Client = &client

	query := `
		SELECT
			a.id,
			a.client_id,
			a.balance,
			a.created_at,
			c.id,
			c.name,
			c.email,
			c.created_at
		FROM
			accounts a
		INNER JOIN
			clients c
		ON c.id = a.client_id
		WHERE
			a.id = ?`

	stmt, err := a.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(
		&account.ID,
		&account.Client.ID,
		&account.Balance,
		&account.CreatedAt,
		&client.ID,
		&client.Name,
		&client.Email,
		&client.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (a *AccountDB) Save(account *entity.Account) error {
	stmt, err := a.DB.Prepare(`
		INSERT INTO accounts (id, client_id, balance, created_at)
		VALUES (?, ?, ?, ?)
	`)
	defer stmt.Close()
	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		account.ID,
		account.Client.ID,
		account.Balance,
		account.CreatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func (a *AccountDB) UpdateBalance(account *entity.Account) error {
	stmt, err := a.DB.Prepare("UPDATE accounts SET balance = ? WHERE id = ?")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(account.Balance, account.ID)
	if err != nil {
		return err
	}

	return nil
}
