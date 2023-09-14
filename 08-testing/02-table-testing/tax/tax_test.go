package tax

import "testing"

func TestCalculateTax(t *testing.T) {
	tests := []struct {
		amount   float64
		expected float64
	}{
		{500.0, 5.0},
		{1000.0, 10.0},
		{1500.0, 10.0},
	}

	for i, test := range tests {
		result := CalculateTax(test.amount)

		if result != test.expected {
			t.Errorf("%d: expected %f, got %f instead.", i, test.expected, result)
		}
	}
}
