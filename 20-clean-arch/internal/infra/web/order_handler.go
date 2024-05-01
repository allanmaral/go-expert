package web

import (
	"encoding/json"
	"net/http"

	"github.com/allanmaral/go-expert/20-clean-arch/internal/usecase"
)

type OrderHandlerWeb struct {
	CreateOrderUseCase *usecase.CreateOrderUseCase
}

func NewOrderHandlerWeb(createOrderUseCase *usecase.CreateOrderUseCase) *OrderHandlerWeb {
	return &OrderHandlerWeb{
		CreateOrderUseCase: createOrderUseCase,
	}
}

func (h *OrderHandlerWeb) Create(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateOrderInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := h.CreateOrderUseCase.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(output); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
}
