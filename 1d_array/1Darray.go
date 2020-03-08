package main

import (
	"fmt"
	"math"
)

func main() {
	arr := [6]int{3, 9, 2, 12, 6, 54}
	printArr(arr[:])
	fmt.Println("Minimum value of array is ", findMinimum(arr[:]))
	fmt.Println("Maximum value of array is ", findMaximum(arr[:]))
	fmt.Printf("Standard deviation of array is %.3f \n", standardDeviation(arr[:]))
	bubbleSort(arr[:])
	printArr(arr[:])
}

func findMinimum(arr []int) int {
	min := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}

func findMaximum(arr []int) int {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func total(arr []int) int {
	total := arr[0]
	for i := 1; i < len(arr); i++ {
		total += arr[i]
	}
	return total
}

func mean(arr []int) float64 {
	return float64(total(arr) / len(arr))
}

func variance(arr []int) float64 {
	mean := mean(arr)
	var sum float64 = 0
	for i := 0; i < len(arr); i++ {
		sum += (float64(arr[i]) - mean) * (float64(arr[i]) - mean)
	}

	return sum / float64(len(arr))
}

func standardDeviation(arr []int) float64 {
	return math.Sqrt(variance(arr))
}

func printArr(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], "\t")
	}
	fmt.Println()
}

func bubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
