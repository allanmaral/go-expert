package entity

import (
	"errors"
	"time"

	"github.com/allanmaral/go-expert/09-apis/pkg/entity"
)

var (
	ErrRequiredID    = errors.New("required id")
	ErrInvalidID     = errors.New("invalid id")
	ErrRequiredName  = errors.New("required name")
	ErrRequiredPrice = errors.New("price is required")
	ErrInvalidPrice  = errors.New("invalid price")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewProduct(name string, price int) (*Product, error) {
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
