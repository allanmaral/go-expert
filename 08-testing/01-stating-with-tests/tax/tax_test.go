package tax

import "testing"

// To run tests, use the command
// `go test .` or `go test -v .` for a verbose run

func TestCalculateTaxWithPassingAssertion(t *testing.T) {
	amount := 500.0
	expected := 5.0

	result := CalculateTax(amount)

	if result != expected {
		t.Errorf("Expected %f, got %f instead", expected, result)
	}
}

func TestCalculateTaxWithFailingAssertion(t *testing.T) {
	amount := 500.0
	expected := 1.0

	result := CalculateTax(amount)

	if result != expected {
		t.Errorf("Expected %f, got %f instead", expected, result)
	}
}
