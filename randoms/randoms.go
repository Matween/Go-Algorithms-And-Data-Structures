package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(randomLCG(0, 5, 3, 7, 10))
	rand.Seed(time.Now().UnixNano())
	var list []int
	list = randomLCG(0, 5, 3, 7, 10)
	randomArr(list, len(list))
	fmt.Println(list)
	people := [...]string{"Matt", "Oleg", "Stefan", "Franz",
		"Ivar", "Loki", "Tom", "Ragnar",
		"Mike", "Annie"}

	winners := randomPartly(people[:], 3)

	fmt.Println("The winners are: ", winners)
}

func randomArr(arr []int, n int) {
	for i := 0; i < n; i++ {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func randomLCG(seed, a, b, m, n int) []int {
	var list []int
	xn := seed
	for i := 0; i < n; i++ {
		xn = (xn*a + b) % m
		list = append(list, xn)
	}

	return list
}

func randomPartly(arr []string, num int) []string {
	for i := 0; i < num; i++ {
		j := rand.Intn(len(arr)-i) + i // rand from i to len(arr)
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr[0:num]
}
