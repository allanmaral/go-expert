package tax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {
	tax, err := CalculateTax(1000.0)

	assert.Nil(t, err)
	assert.Equal(t, 10.0, tax)
}

func TestCalculateTaxShouldReturnErrorOnZeroAmount(t *testing.T) {
	tax, err := CalculateTax(0.0)

	assert.Error(t, err, "amount must be greater then 0")
	assert.Equal(t, 2.0, tax)
}
