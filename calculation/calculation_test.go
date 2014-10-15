package calculation

import (
	"testing"
	"github.com/flower-pot/primes/calculation"
)

func TestCalculatePrimes(t *testing.T) {
	calculation.CalculatePrimes(2, 100000)
}