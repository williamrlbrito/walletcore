package web

import (
	"encoding/json"
	"net/http"

	"github.com/williamrlbrito/walletcore/internal/usecase/create_transaction"
)

type WebTransactionHandler struct {
	CreateTransactionUseCase create_transaction.CreateTransactionUseCase
}

func NewWebTransactionHandler(createTransactionUseCase create_transaction.CreateTransactionUseCase) *WebTransactionHandler {
	return &WebTransactionHandler{CreateTransactionUseCase: createTransactionUseCase}
}

func (handler *WebTransactionHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var dto create_transaction.CreateTransactionInputDTO

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	output, err := handler.CreateTransactionUseCase.Execute(ctx, dto)
	if err != nil {
		println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
