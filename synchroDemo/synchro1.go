package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
atomic operations is one method of achieving synchronization across goroutines.
use Add and Load methods from sync.atomic pkg.
*/

func incrAtomicCntr(c *atomic.Uint64) {
	for range 5000 {
		c.Add(1)
	}
}

func atomicOpsDemo() {
	var cntr atomic.Uint64
	var wg sync.WaitGroup

	for range 5 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			incrAtomicCntr(&cntr)
		}()
	}
	wg.Wait()
	fmt.Println("atomicDemo cntr = ", cntr.Load())
}

func incrCntr(c *uint64) {
	for range 5000 {
		(*c)++ //:)
	}
}
func mutexDemo() {
	var cntr uint64
	var wg sync.WaitGroup
	var mtx sync.Mutex

	for range 5 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer mtx.Unlock()
			mtx.Lock()
			incrCntr(&cntr)
		}()
	}
	wg.Wait()
	fmt.Println("mutexDemo cntr = ", cntr)
}

func main() {
	atomicOpsDemo()
	mutexDemo()
}
