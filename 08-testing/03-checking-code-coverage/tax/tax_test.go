package tax

import "testing"

// Code coverage can be evaluated using the flag `-coverprofile=<coverage-output-file>`
// Eg:
//	go test -coverprofile=coverage.out .
//
// We can use the `go tool cover -html=coverage.out` to inspect the coverage result

func TestCalculateTax(t *testing.T) {
	tests := []struct {
		amount   float64
		expected float64
	}{
		{500.0, 5.0},
		{1000.0, 10.0},
		{1500.0, 10.0},
		{0.0, 0.0},
	}

	for i, test := range tests {
		result := CalculateTax(test.amount)

		if result != test.expected {
			t.Errorf("%d: expected %f, got %f instead.", i, test.expected, result)
		}
	}
}
