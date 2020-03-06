package main

import "fmt"

func main() {
	fmt.Println("Greatest common divisor: ", gcd(11, 66))
	fmt.Println("Least common multiple: ", lcm(100000, 200000))
}

func gcd(a, b int) int {
	for b != 0 {
		remainder := a % b
		a = b
		b = remainder
	}
	return a
}

func lcm(a, b int) int {
	return (a / gcd(a, b)) * b
}
