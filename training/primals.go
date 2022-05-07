package main

import (
	"fmt"
	"math"
)

func is_prime(number float64) bool {
	if number == 2 {
		return true
	}
	for i := 2; i < int(math.Sqrt(number)+1); i++ {
		if int(number)%i == 0 {
			return false
		}
	}
	return true
}

func primals(number int) []int {
	primes := make([]int, 0)
	for i := 2; i < number+1; i++ {
		if is_prime(float64(i)) {
			primes = append(primes, i)

		}
	}
	return primes
}

func main() {
	fmt.Println(is_prime(7))
	fmt.Println(primals(10))
}
