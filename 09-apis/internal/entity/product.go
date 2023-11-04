package entity

import (
	"errors"
	"time"

	"github.com/allanmaral/go-expert/09-apis/pkg/entity"
)

var (
	ErrRequiredName  = errors.New("required name")
	ErrRequiredPrice = errors.New("price is required")
	ErrInvalidPrice  = errors.New("invalid price")
)

type Product struct {
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	ID        entity.ID `json:"id"`
}

func NewProduct(name string, price float64) (*Product, error) {
	if name == "" {
		return nil, ErrRequiredName
	}
	if price == 0 {
		return nil, ErrRequiredPrice
	}
	if price < 0 {
		return nil, ErrInvalidPrice
	}

	return &Product{
		ID:        entity.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}, nil
}
