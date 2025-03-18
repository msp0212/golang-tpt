package main

import (
	"fmt"
)

type emp struct {
	id int
	name string
}

// func printDetails has a receiver of type emp
// which means it can be invoked via an object of type emp
// by pointer and by value both are supported
func (e *emp) printDetails() {
	fmt.Println("employee details: id = ", e.id, "name = ", e.name)
}

func main() {
	e1 := emp{101, "emp101"}

	ep := &e1

	// conversion between pointer and value types is automatically handled for
	// method calls
	e1.printDetails()
	ep.printDetails()
}

