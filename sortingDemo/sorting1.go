package main

import (
	"cmp"
	"fmt"
	"slices"
)

/*
Sort() from slices pkg is generic and works for ordered built-in types
like string, integers, floats etc
*/
func sliceSortDemo() {
	fmt.Println("sliceSortDemo...")

	fruits := []string{"mango", "apple", "kiwi", "apricot", "plum"}
	fmt.Println("fruits = ", fruits)
	slices.Sort(fruits)
	fmt.Println("sorted = ", fruits)

	tempC := []int{25, 0, 37, 5, -5, -20}
	fmt.Println("temperatures = ", tempC)
	slices.Sort(tempC)
	fmt.Println("sorted temperatures= ", tempC)
	s := slices.IsSorted(tempC)
	fmt.Println("Are temperaturs sorted?", s)
}

/*
SortFunc() from slices pkg can be used to sort elements
by providing a custom sort function.
This is useful if we want to sort some data by something other than its natural
order.
It is also useful of sorting user defined data types
*/

func sliceSortFuncDemo() {
	fmt.Println("sliceSortFuncDemo...")

	// sort fruits by length of their names
	// which is different from the natural alphabetical ordering
	fruits := []string{"mango", "apple", "kiwi", "apricot", "plum"}
	fmt.Println("fruits = ", fruits)
	sortFruitsFunc := func(f1, f2 string) int {
		return cmp.Compare(len(f1), len(f2))
	}
	slices.SortFunc(fruits, sortFruitsFunc)
	fmt.Println("sorted = ", fruits)

	type employee struct {
		name string
		age  uint
	}
	emps := []employee{
		{"emp1", 35},
		{"emp2", 30},
		{"emp3", 30},
		{"emp4", 45},
		{"emp5", 25},
	}
	fmt.Println("emps = ", emps)
	sortEmployeesFunc := func(e1, e2 employee) int {
		return cmp.Compare(e1.age, e2.age)
	}
	slices.SortFunc(emps, sortEmployeesFunc)
	fmt.Println("sorted emps = ", emps)
}

func main() {
	sliceSortDemo()
	sliceSortFuncDemo()
}
