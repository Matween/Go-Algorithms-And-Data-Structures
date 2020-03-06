package main

import (
	"fmt"
	"math"
)

func main() {
	printFactorization(primeFactorization(196))
}

func primeFactorization(number int) []int {
	var factors []int
	for number%2 == 0 {
		factors = append(factors, 2)
		number = number / 2
	}
	for i := 3; i < int(math.Sqrt(float64(number)))+1; i += 2 {
		for number%i == 0 {
			factors = append(factors, i)
			number = number / i
		}
	}
	if number > 2 {
		factors = append(factors, number)
	}
	return factors
}

func printFactorization(factors []int) {
	for i, factor := range factors {
		if i != len(factors)-1 {
			fmt.Print(factor, "x")
		} else {
			fmt.Print(factor)
		}
	}
}
