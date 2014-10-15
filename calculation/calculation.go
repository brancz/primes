package calculation

import (
	"fmt"
	"errors"
	"github.com/flower-pot/primes/algorithm"
)

const no_prime int = -1

func CalculatePrimes(start, end int) ([]int, error) {
	if start < 0 || end < 0 || start > end {
		return nil, errors.New("Start and end values are invalid")
	}
	if start % 2 == 0 { //Start is an even number
		start ++	//So skip it, causes skipping of other even numbers
	}
	list := make([]int, 0)
	c := make(chan int)
	for i := start; i < end; i += 2 {
		go checkPrime(c, i)
	}
	for i := start; i < end; i += 2 {
		if n := <- c; n != no_prime {
			list = append(list, n)
			fmt.Println(n)
		}
	}
	return list, nil
}

func checkPrime(channel chan int, in int) {
	if algorithm.IsPrime(in) {
		channel <- in
		return
	}
	channel <- -1
	return
}