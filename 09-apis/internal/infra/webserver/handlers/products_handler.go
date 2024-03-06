package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/allanmaral/go-expert/09-apis/internal/dto"
	"github.com/allanmaral/go-expert/09-apis/internal/entity"
	"github.com/allanmaral/go-expert/09-apis/internal/infra/database"
	"github.com/go-chi/chi/v5"
)

type ProductHandler struct {
	ProductRepository database.ProductRepository
}

func NewProductHandler(db database.ProductRepository) *ProductHandler {
	return &ProductHandler{
		ProductRepository: db,
	}
}

// Create Product godoc
// @Summary		Create product
// @Description	Create a product
// @Tags		products
// @Accept		json
// @Product		json
// @Param		request		body		dto.CreateProductInput	true	"Product data"
// @Success		201
// @Failure		401			{object}	dto.ErrorOutput
// @Failure		403			{object}	dto.ErrorOutput
// @Failure		500			{object}	dto.ErrorOutput
// @Router		/products [post]
// @Security	ApiKeyAuth
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

	w.Header().Set("Location", fmt.Sprintf("/products/%s", p.ID))
	w.WriteHeader(http.StatusCreated)
}

// List Product godoc
// @Summary		List products
// @Description	List all products
// @Tags		products
// @Accept		json
// @Product		json
// @Param		page		query		string					false	"Page number"
// @Param		limit		query		string					false	"Page page size limit"
// @Success		200			{array}		dto.ProductOutput
// @Failure		401			{object}	dto.ErrorOutput
// @Failure		403			{object}	dto.ErrorOutput
// @Failure		500			{object}	dto.ErrorOutput
// @Router		/products [get]
// @Security	ApiKeyAuth
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	sort := r.URL.Query().Get("sort")

	pagenum, err := strconv.Atoi(page)
	if err != nil {
		pagenum = 0
	}

	limitnum, err := strconv.Atoi(limit)
	if err != nil {
		limitnum = 0
	}

	products, err := h.ProductRepository.FindAll(pagenum, limitnum, sort)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	output := make([]dto.ProductOutput, len(products))
	for i, p := range products {
		output[i] = dto.ProductOutput{ID: p.ID.String(), Name: p.Name, Price: p.Price}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

// Get Product godoc
// @Summary		Get a product
// @Description	Get a specific product
// @Tags		products
// @Accept		json
// @Product		json
// @Param		id			path		string					true	"Product Identifier"	Format(uuid)
// @Success		200			{object}	dto.ProductOutput
// @Failure		401			{object}	dto.ErrorOutput
// @Failure		403			{object}	dto.ErrorOutput
// @Failure		404			{object}	dto.ErrorOutput
// @Failure		500			{object}	dto.ErrorOutput
// @Router		/products/{id} [get]
// @Security	ApiKeyAuth
func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := h.ProductRepository.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	p := dto.ProductOutput{
		ID:    product.ID.String(),
		Name:  product.Name,
		Price: product.Price,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(p)
}

// Update Product godoc
// @Summary		Update a product
// @Description	Update a specific product
// @Tags		products
// @Accept		json
// @Product		json
// @Param		id			path		string					true	"Product Identifier"	Format(uuid)
// @Param		request		body		dto.UpdateProductInput	true	"Updated product data"
// @Success		200			{object}	dto.ProductOutput
// @Failure		401			{object}	dto.ErrorOutput
// @Failure		403			{object}	dto.ErrorOutput
// @Failure		404			{object}	dto.ErrorOutput
// @Failure		500			{object}	dto.ErrorOutput
// @Router		/products/{id} [put]
// @Security	ApiKeyAuth
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	product, err := h.ProductRepository.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var in dto.UpdateProductInput
	err = json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product.Name = in.Name
	product.Price = in.Price
	err = h.ProductRepository.Update(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	output := dto.ProductOutput{
		ID:    product.ID.String(),
		Name:  product.Name,
		Price: product.Price,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

// Delete Product godoc
// @Summary		Delete a product
// @Description	Delete a specific product
// @Tags		products
// @Accept		json
// @Product		json
// @Param		id			path		string					true	"Product Identifier"	Format(uuid)
// @Success		204			{object}	dto.ProductOutput
// @Failure		401			{object}	dto.ErrorOutput
// @Failure		403			{object}	dto.ErrorOutput
// @Failure		404			{object}	dto.ErrorOutput
// @Failure		500			{object}	dto.ErrorOutput
// @Router		/products/{id} [delete]
// @Security	ApiKeyAuth
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	_, err := h.ProductRepository.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = h.ProductRepository.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
