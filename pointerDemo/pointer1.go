package main

import (
	"fmt"
)

// pass by value
func func1(num int) {
	num = 100;
}

// pass by reference or address
func func2(num *int) {
	*num = 100
}

func main() {
	fmt.Println("Pointer Demo")

	n := 10
	fmt.Println("n = ", n)

	func1(n)
	fmt.Println("n after func1 call= ", n)

	func2(&n)
	fmt.Println("n after func2 call ", n)
}
