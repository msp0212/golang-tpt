package main

import (
	"fmt"
	"strconv"
)

func printSlice(s []string, prefix string) {
	fmt.Print(prefix, " ")
	fmt.Println("len = ", len(s), "cap = ", cap(s), "slice = ", s)
}

func main() {
	var s1 []string
	printSlice(s1, "uninit'd slice s1")

	s2 := make([]string, 0, 5)
	printSlice(s2, "slice s2 created by make")

	for i := 0; i < 5; i++ {
		s2 = append(s2, "str"+strconv.Itoa(i))
	}
	printSlice(s2, "slice s2 after appending 5 value")

	s2 = append(s2, "str5")
	printSlice(s2, "slice s2 after appending 6th value")

	for i := 6; i < 10; i++ {
		s2 = append(s2, "str"+strconv.Itoa(i))
	}
	printSlice(s2, "slice s2 after appending 4 more values")

	s2 = append(s2, "str11")
	printSlice(s2, "slice s2 after appending 11th value")

	sl1 := s2[2:4]
	printSlice(sl1, "slice of s2 [2:4]")

	sl2 := make([]string, len(sl1), cap(sl1))
	n := copy(sl2, sl1)
	fmt.Println("copied ", n, "elems from sl1 to sl2")
	printSlice(sl2, "sl2 copied from sl1")

	sl2 = append(sl2, sl1...)
	printSlice(sl2, "sl2 after appending sl1 to it")
}
