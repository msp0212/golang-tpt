package main

import "fmt"

var pf = fmt.Printf

func printFormattingDemo() {
	fmt.Println("printFormattingDemo...")

	type employee struct {
		name string
		age  uint
	}
	e1 := employee{"emp1", 25}

	//struct printing
	pf("struct emp: %v\n", e1)
	pf("struct emp: %+v\n", e1)
	pf("struct emp: %#v\n", e1)

	//type printing
	pf("type emp: %T\n", e1)

	//pointer printing
	pf("pointer: %p\n", &e1)

	//bool printing
	pf("bool: %t\n", true)
	pf("bool: %t\n", false)

	//integer printing
	pf("int: %d\n", 14) //base 10
	pf("int: %b\n", 14) // base 2
	pf("int: %x\n", 14) // base 16
	pf("int: %X\n", 14) // base 16
	pf("int: %c\n", 48) // char for given int val
	pf("int with width: |%3d|%3d|%3d|\n", 1, 11, 111)
	pf("int with width left justified: |%-3d|%-3d|%-3d|\n", 1, 11, 111)

	//float printing
	pf("float: %f\n", 123.45)
	pf("float: %e\n", 12345678.99)
	pf("float: %E\n", 12345678.99)
	pf("float with width and precision: |%6.2f|%6.2f|%6.2f|\n", 1.1, 1.11, 1.111)
	pf("float with width and precision left justified: |%-6.2f|%-6.2f|%-6.2f|\n", 1.1, 1.11, 1.111)

	//string printing
	pf("string: %s\n", "\"foobar\"")
	pf("string: %q\n", "\"foobar\"")
	pf("string: %x\n", "ABC123")
	pf("string with width: |%10s|%10s|\n", "foo", "foobar")
	pf("string with width left justified: |%-10s|%-10s|\n", "foo", "foobar")

}

func main() {
	printFormattingDemo()
}
