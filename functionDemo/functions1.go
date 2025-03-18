package main

import (
	"fmt"
)

// function with single return value
func sum1(a, b int) int {
	return a + b
}

// function with multiple return values
func sum2(a, b int) (int, bool) {
	sum := a + b
	var err = false

	if sum > 100 {
		err = true
	}
	return sum, err
}

// variadic function
// variable args are in form of slice
func sum3(nums ...int) int {

	fmt.Println("slice of len ", len(nums), "recvd in func params = ", nums)
	sum := 0

	for _, n := range nums {
		sum += n
	}
	return sum
}

func main() {

	s := sum1(1, 2)
	fmt.Println("sum1 of 1 and 2 = ", s)

	s, e1 := sum2(60, 70)
	fmt.Println("sum2 of 60 and 70 = ", s, ", error = ", e1)

	s = sum3(1, 2, 3, 4, 5)
	fmt.Println("sum3 of 1 to 5 = ", s)

	nums := []int{1, 2, 3, 4, 5, 6}
	s = sum3(nums...)
	fmt.Println("sum3 of 1 to 6 = ", s)
}
