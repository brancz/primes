package calculation

import (
	"testing"
	"github.com/flower-pot/primes/calculation"
)

func TestCalculatePrimes(t *testing.T) {
	expectPrimes(0, 5, []int{2, 3, 5}, t)
	expectPrimes(2, 5, []int{2, 3, 5}, t)
	expectPrimes(5, 10, []int{5, 7}, t)
}

func expectPrimes(from int, to int, primes []int, t *testing.T) {
	list, err := calculation.CalculatePrimes(from, to)
	if err != nil {
		t.Errorf("An error occured")
		return
	}
	if !equal(list, primes) {
		t.Errorf("%v has not all expected values", list)
	}
}

func equal(a, b []int) bool {
    if len(a) != len(b) {
        return false
    }

    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }

    return true
}