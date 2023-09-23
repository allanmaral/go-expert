package entity

import (
	"testing"

	generators "github.com/allanmaral/go-expert/09-apis/test"
	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	name := generators.RandonName()
	price := generators.RandonPrice()

	p, err := NewProduct(name, price)

	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.NotEmpty(t, p.ID)
	assert.Equal(t, name, p.Name)
	assert.Equal(t, price, p.Price)
}

func TestNewProductShouldReturnInvalidPriceErrorOnNegativePrice(t *testing.T) {
	tests := []int64{-1, -2, -10, -99, -100, -101, -999, -1000, -1001}

	for _, test := range tests {
		p, err := NewProduct("any name", test)

		assert.Nil(t, p)
		assert.Equal(t, ErrInvalidPrice, err)
	}
}

func TestNewProductShouldReturnRequiredPriceErrorOnZeroPrice(t *testing.T) {
	p, err := NewProduct("any name", 0)

	assert.Nil(t, p)
	assert.Equal(t, ErrRequiredPrice, err)
}

func TestNewProductShouldReturnRequiredNameErrorOnEmptyName(t *testing.T) {
	emptyName := ""
	p, err := NewProduct(emptyName, generators.RandonPrice())

	assert.Nil(t, p)
	assert.Equal(t, ErrRequiredName, err)
}
