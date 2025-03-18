package main

import (
	"fmt"
)

// introduce a new type "emp"
type emp struct {
	name string
	id   int
}

// func to create an object of type emp
// it is ok to return address of local var since go is gc'd lang
func newEmp(name string, id int) *emp {
	e := emp{name: name, id: id}
	return &e
}

func main() {

	// new obj of type emp
	e1 := emp{"emp1", 1}
	fmt.Println("e1 = ", e1)

	// new obj can be initialised by explicitly naming the fields
	e2 := emp{name: "emp2", id: 2}
	fmt.Println("e2 = ", e2)

	// it is ok to not give initialisation for few fields, they will zero'd out
	e3 := emp{name: "emp3"}
	e4 := emp{id: 4}
	fmt.Println("e3 = ", e3)
	fmt.Println("e4 = ", e4)

	//good practice to have func for creating objects of a given type
	e5 := newEmp("emp5", 5)
	fmt.Println("e5 = ", e5)

	//fields can be accesed individually by their name using . operator
	fmt.Println("e5.name = ", e5.name, " e5.id = ", e5.id)

	//anonymous struct
	addr1 := struct {
		pincode int
		city string
	} {
		pincode: 123456,
		city: "abc",
	}

	fmt.Println("addr1 = ", addr1)
}
