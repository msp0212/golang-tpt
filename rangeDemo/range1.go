package main

import (
	"fmt"
)

func main() {
	fmt.Println("Range Demo")

	// range on slice
	nums := []int{1, 2, 3, 4, 5}
	for inx, val := range nums {
		fmt.Println("val at index ", inx, "=", val)
	}

	// range on map
	capitals := map[string]string{"india": "new delhi", "nepal": "kathmandu"}
	for key, val := range capitals {
		fmt.Println("Country : ", key, ", Capital: ", val)
	}

	// range on strings
	for inx, val := range "ABCD" {
		fmt.Println("val at index ", inx, "=", val)
	}
}
