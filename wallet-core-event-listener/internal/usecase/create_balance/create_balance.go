package create_balance

import (
	"github.com/danyukod/wallet-core-event-listener/internal/entity"
	"github.com/danyukod/wallet-core-event-listener/internal/gateway"
)

type CreateBalanceInputDTO struct {
	AccountID string  `json:"account_id"`
	Amount    float64 `json:"amount"`
}

type CreateBalanceUseCase struct {
	gateway.BalanceGateway
}

func NewCreateBalanceUseCase(balanceGateway gateway.BalanceGateway) *CreateBalanceUseCase {
	return &CreateBalanceUseCase{
		BalanceGateway: balanceGateway,
	}
}

func (c CreateBalanceUseCase) Execute(dto CreateBalanceInputDTO) error {
	balance, err := entity.NewBalance(dto.AccountID, dto.Amount)
	if err != nil {
		return err
	}
	err = c.BalanceGateway.Save(balance)
	if err != nil {
		return err
	}

	return nil
}
