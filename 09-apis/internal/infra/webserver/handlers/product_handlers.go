package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/allanmaral/go-expert/09-apis/internal/dto"
	"github.com/allanmaral/go-expert/09-apis/internal/entity"
	"github.com/allanmaral/go-expert/09-apis/internal/infra/database"
)

type ProductHandler struct {
	ProductRepository database.ProductRepository
}

func NewProductHandler(db database.ProductRepository) *ProductHandler {
	return &ProductHandler{
		ProductRepository: db,
	}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	p, err := entity.NewProduct(product.Name, product.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.ProductRepository.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
