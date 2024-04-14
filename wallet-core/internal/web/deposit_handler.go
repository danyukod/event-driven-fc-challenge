package web

import (
	"encoding/json"
	"fmt"
	"github.com.br/devfullcycle/fc-ms-wallet/internal/usecase/create_deposit"
	"net/http"
)

type DepositHandler struct {
	CreateDepositUseCase create_deposit.CreateDepositUseCase
}

func NewDepositHandler(depositUseCase create_deposit.CreateDepositUseCase) *DepositHandler {
	return &DepositHandler{
		CreateDepositUseCase: depositUseCase,
	}
}

func (h *DepositHandler) CreateDeposit(w http.ResponseWriter, r *http.Request) {
	var dto create_deposit.CreateDepositInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	err = h.CreateDepositUseCase.Execute(dto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
