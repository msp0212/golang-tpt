package main

import (
	"fmt"
	"time"
)

/*
goroutines are lightweight threads of execution managed by go runtime.
goroutines run in the same address space, so access to shared resources
must be synchronised.
*/

func printWord(word string, times int) {
	for i := range times {
		fmt.Printf("%s: %d\n", word, i);
	}
}

func main() {
	// synchronous function call
	printWord("sycnhronous func call", 3);

	// evaluation of printWord and its arg happens in current goroutine
	// execution of printWord happens concurrently in a new goroutine
	go printWord("func inside a goroutine", 5);

	// execute an anonymous function inside a goroutine
	go func(word string, times int) {
		for i := range times {
			fmt.Printf("%s: %d\n", word, i);
		}
	} ("anonymous func call inside a goroutine", 5);

	// sleep for 1 second
	time.Sleep(time.Second * 1)
	fmt.Println("Bye");
}
