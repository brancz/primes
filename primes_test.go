package primes

import "github.com/flower-pot/primes"
import "testing"

func TestCalculatePrimes(t *testing.T) {
	primes.CalculatePrimes(2, 100000)
}