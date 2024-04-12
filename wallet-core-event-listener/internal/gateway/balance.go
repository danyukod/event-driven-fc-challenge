package gateway

import "github.com/danyukod/wallet-core-event-listener/internal/entity"

type BalanceGateway interface {
	FindByAccountID(accountID string) (*entity.Balance, error)
	Save(*entity.Balance) error
}
