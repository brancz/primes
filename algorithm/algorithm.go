package algorithm

import (
	"math"
)

func IsPrime(n int) bool {
	if n < 2 || n == 4 || n == 6 {
		return false
	}
	if n == 2 {
		return true
	}
	end := int(math.Ceil(math.Sqrt(float64(n))))
	for i := 3; i < end; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
