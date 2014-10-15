package algorithm

import (
	//"fmt"
	//"math/big"
)

func IsPrime(n int) bool {
	if n < 2 || n == 4 {
		return false
	}
	if n == 2{
		return true
	}
	for i := 3; i < n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}