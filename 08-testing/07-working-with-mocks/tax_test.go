package tax

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
)

type TaxRepositoryMock struct {
	mock.Mock
}

func (m *TaxRepositoryMock) SaveTax(tax float64) error {
	args := m.Called(tax)
	return args.Error(0)
}

var _ Repository = (*TaxRepositoryMock)(nil)

func TestCalculateTaxShouldNotReturnErrorOnNonZeroValue(t *testing.T) {
	repository := &TaxRepositoryMock{}
	repository.On("SaveTax", 10.0).Return(nil).Once()

	err := CalculateTaxAndSave(1000.0, repository)

	repository.AssertExpectations(t)
	if err != nil {
		t.Errorf("expected nil, got %v instead", err)
	}
}

func TestCalculateTaxShouldReturnErrorOnZeroValue(t *testing.T) {
	repository := &TaxRepositoryMock{}
	repository.On("SaveTax", 0.0).Return(errors.New("error saving tax")).Once()

	err := CalculateTaxAndSave(0.0, repository)

	repository.AssertExpectations(t)
	if err == nil {
		t.Errorf("expected \"error saving tax\", got %v instead", err)
	}
}
