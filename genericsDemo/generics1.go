package main

import (
	"fmt"
)

func intSum(s []int64) int64 {
	var sum int64
	sum = 0
	for _, v := range s {
		sum += v
	}
	return sum
}

func floatSum(s []float64) float64 {
	sum := 0.0
	for _, v := range s {
		sum += v
	}
	return sum
}

// Generic function definition with type information provided in square brackets
// refer https://go.dev/doc/tutorial/generics
func genericSum[T int64 | float64](s []T) T {
	var sum T
	for _, v := range s {
		sum += v
	}
	return sum
}

func main() {
	intSlice := []int64{1, 2, 3, 4, 5}
	floatSlice := []float64{1.1, 2.2, 3.3, 4.4, 5.5}

	fmt.Println("Int Sum:", intSum(intSlice))
	fmt.Println("Float Sum:", floatSum(floatSlice))

	// Generic function call, type information provided in square brackets
	fmt.Println("Generic Int Sum:", genericSum(intSlice))
	fmt.Println("Generic Int Sum:", genericSum[int64](intSlice))
	fmt.Println("Generic Float Sum:", genericSum(floatSlice))
	fmt.Println("Generic Float Sum:", genericSum[float64](floatSlice))
}
