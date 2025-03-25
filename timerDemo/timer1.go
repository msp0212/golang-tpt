package main

import (
	"fmt"
	"time"
)

/*
Timers:
time.NewTimer creates a new timer(one time).
Represents a single event in future
Tickers:
time.NewTicker creatse a new ticker(recurring).
Represent a periodic event in future
*/

func startTimer(timeoutSec int)
{
	fmt.Println("Will wake you up after", timeoutSec, "seconds")
	timer := time.NewTimer(time.Second * time.Duration(timeoutSec))

	go func() {
		select {
		case t := <- timer.C:
			return
		}
	}
	fmt.Println("Waking you up after", timeoutSec, "seconds at", t)
}

// wake me up after specified seconds
func wakeMeUp(timeoutSec int) {
	fmt.Println("===Timer Demo===")
	startTimer(1)
	startTimer(5)
	fmt.Println("===Timer Demo done===")
}

// ping me periodically
func pingMePeriodically(timeoutSec int) {
	ticker := time.NewTicker(time.Second * time.Duration(timeoutSec))
	for {
		select {
		case t := <- ticker.C:
			fmt.Println("Pinging you up after", timeoutSec, "seconds at", t)
		case t := <- time.After(time.Second * 10):
			fmt.Println("Stopping the pings now at", t)
			return
		default:
		}
	}
}

func main() {
	wakeMeUp(2)
	pingMePeriodically(1)
}
