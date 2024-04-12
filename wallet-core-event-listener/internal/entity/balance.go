package entity

import (
	"errors"
	"github.com/google/uuid"
)

type Balance struct {
	ID        string  `json:"id"`
	AccountID string  `json:"account_id"`
	Amount    float64 `json:"amount"`
}

func NewBalance(accountID string, amount float64) (*Balance, error) {
	balanceEntity := &Balance{
		ID:        uuid.New().String(),
		AccountID: accountID,
		Amount:    amount,
	}
	err := balanceEntity.Validate()
	if err != nil {
		return nil, err
	}
	return balanceEntity, nil
}

func (b Balance) Validate() error {
	if b.AccountID == "" {
		return errors.New("account_id is required")
	}
	return nil
}
