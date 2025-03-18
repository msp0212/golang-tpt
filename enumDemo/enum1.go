package main

import (
	"fmt"
)

/*
Go doesn't provide explicit enum type
But it can be modelled using existing go idioms
*/

type EmpType int

//enumerate the possible list of values
const (
	FullTime EmpType = iota // iota keyword generates successive values 0,1,2...
	PartTime
	Contracter
)

//provide a map for value to string conversion
var EmpTypeStrMap = map[EmpType]string {
	FullTime: "Full-Time",
	PartTime: "Part-Time",
	Contracter: "Contracter",
}

//implement Stringer interface for converting value to string format
func (et EmpType) String() string {
	return EmpTypeStrMap[et]
}


func main() {
	et1 := FullTime

	fmt.Println("et1 = ", et1);

	var et2 EmpType; // default value is zero-value
	fmt.Println("et2 = ", et2);

	var et3 EmpType;
	// et3 = 2; allowed since 2 is valid value i.e. Contracter
	// et3 = 4;  not allowed
	et3 = Contracter;
	fmt.Println("et3 = ", et3);
}
