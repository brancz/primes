package algorithm

import (
	"testing"
	"github.com/flower-pot/primes/algorithm"
)

func TestIsPrime(t *testing.T) {
	actualPrimes := []int{2, 3, 5, 7, 11, 13, 17, 19, 524287}
	for _, n := range actualPrimes {
		expectPrime(n, t)
	}

	notPrimes := []int{1, 0, -1, 6, 4, 100, 1000, 100000}
	for _, n := range notPrimes {
		expectNotPrime(n, t)
	}
}

func expectPrime(in int, t *testing.T) {
	if val := algorithm.IsPrime(in); val != true {
		t.Errorf("%v is a prime number, but function said %v", in, val)
	}
}

func expectNotPrime(in int, t *testing.T) {
	if val := algorithm.IsPrime(in); val == true {
		t.Errorf("%v is not a prime number, but function said %v", in, val)
	}
}