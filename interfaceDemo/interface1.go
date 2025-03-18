package main

import (
	"fmt"
)


// an interface is a named collection of method signatures

type animal interface {
	speak() string
	move() string
}

func interact(a animal) {
	//an oject of interface type can call all the methods in that interface
	s := a.speak()
	m := a.move()

	//get type of the object which called the interface method at runtime
	var at string
	switch a.(type) {
	case sparrow:
		at = "sparrow"
	case dog:
		at = "dog"
	default:
		at = "unknown"
	}
	fmt.Println("Animal = ", at, ", speak = ", s, "and move = ", m)
}


// struct sparrow implementing methods of animal interface
type sparrow struct {
	color string
}

func (s sparrow) speak() string {
	return "chirp"
}

func (s sparrow) move() string {
	return "fly"
}


// struct dog implementing methods of animal interface
type dog struct {
	breed string;
	
}

func (s dog) speak() string {
	return "bark"
}

func (s dog) move() string {
	return "walk"
}

func main() {
	s :=  sparrow{}
	d := dog{}
	// any object type implementing all the interface methods can be used
	// for an interface type argument
	interact(s)
	interact(d)
}
