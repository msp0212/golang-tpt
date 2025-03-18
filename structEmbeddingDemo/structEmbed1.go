package main

import (
	"fmt"
)

type personId struct {
	name string
	uid  uint64
}

func (p personId) printDetails() {
	fmt.Println("Name:", p.name)
	fmt.Println("UID:", p.uid)
}

type person struct { // struct embedding
	personId   // person struct has a field of type personId, personId is embedded in person struct
	occupation string
}

func main() {
	p := person{
		personId: personId{
			name: "John Doe",
			uid:  123456,
		},
		occupation: "Software Engineer", // occupation is a field of person struct
	} // p is a struct of type person

	// Accessing fields of person struct
	fmt.Println("Person Details:", p.name, p.uid, p.occupation)
	// Accessing base struct fields using base struct name
	fmt.Println("Person Details:", p.personId.name, p.personId.uid, p.occupation)
	// base method can be called using outer object
	p.printDetails() // calling printDetails method of personId struct

	type printDetailsInterface interface {
		printDetails()
	}

	// Interface type variable can hold the object of struct which implements the interface
	var pdi printDetailsInterface = p
	pdi.printDetails()
}
