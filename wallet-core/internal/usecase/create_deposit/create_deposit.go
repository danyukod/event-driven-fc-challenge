package create_deposit

import (
	"errors"
	"github.com.br/devfullcycle/fc-ms-wallet/internal/gateway"
)

type CreateDepositInputDTO struct {
	AccountID string  `json:"account_id"`
	Amount    float64 `json:"amount"`
}

type CreateDepositUseCase struct {
	gateway.AccountGateway
}

func NewCreateDepositUseCase(a gateway.AccountGateway) *CreateDepositUseCase {
	return &CreateDepositUseCase{
		AccountGateway: a,
	}
}

func (uc *CreateDepositUseCase) Execute(input CreateDepositInputDTO) error {
	account, err := uc.AccountGateway.FindByID(input.AccountID)
	if err != nil {
		return err
	}
	if input.Amount <= 0 {
		return errors.New("amount must be greater than 0")
	}
	newBalance := account.Balance + input.Amount
	account.Balance = newBalance
	err = uc.AccountGateway.UpdateBalance(account)
	if err != nil {
		return err
	}
	return nil
}
