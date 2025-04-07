package main

import (
	"fmt"
	"sync"
	"time"
)

func doSomeWork(job string) {
	fmt.Println("Start working on", job)
	time.Sleep(time.Second)
	fmt.Println("Stop working on", job)
}

/*
waitgroup is mechanism to let a function wait for end of all the goroutines
spawned from it.
waitgroup Add() increments the number of goroutines a function is waiting for.
waitgroup Done() decrement the number of goroutines a function is waiting for
waitgroup Wait() block till number of goroutines being waited becomes zero.
*/

func main() {
	var jobs = []string{
		"fetch-water",
		"fetch-vegetables",
		"fetch-rice",
		"fetch-sugar",
		"fetch-milk",
	}
	var wg sync.WaitGroup
	for _, job := range jobs {
		wg.Add(1)
		go func(job string) {
			/*
				defer statement delays the execution of function supplied to it
				after the return of the enclosing function.
				Note: func argument evaluation is not delayed,
						only the func call is delayed
			*/
			defer wg.Done()
			doSomeWork(job)
		}(job)
	}
	wg.Wait()
}
