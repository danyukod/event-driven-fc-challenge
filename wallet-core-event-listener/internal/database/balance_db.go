package database

import (
	"database/sql"
	"github.com/danyukod/wallet-core-event-listener/internal/entity"
)

type BalanceDB struct {
	DB *sql.DB
}

func NewBalanceDB(db *sql.DB) *BalanceDB {
	return &BalanceDB{
		DB: db,
	}
}

func (b *BalanceDB) FindByAccountID(accountID string) (*entity.Balance, error) {
	var balance entity.Balance

	stmt, err := b.DB.Prepare("SELECT b.id, b.account_id, b.balance FROM balance b WHERE b.account_id = ? ORDER BY created_at DESC LIMIT 1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(accountID)
	err = row.Scan(&balance.ID, &balance.AccountID, &balance.Amount)
	if err != nil {
		return nil, err
	}
	return &balance, nil
}

func (b *BalanceDB) Save(balance *entity.Balance) error {
	stmt, err := b.DB.Prepare("INSERT INTO balance (id, account_id, balance) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(balance.ID, balance.AccountID, balance.Amount)
	if err != nil {
		return err
	}
	return nil
}
