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

func startTimer(timeoutSec int) {
	fmt.Println("Will wake you up after", timeoutSec, "seconds")
	timer := time.NewTimer(time.Second * time.Duration(timeoutSec))

	go func() {
		select {
		case t := <- timer.C:
			fmt.Println("Waking you up after", timeoutSec, "seconds at", t)
			return
		case <- time.After(time.Second * 3):
			// all timers > 3 secs will not be run :)
			stop := timer.Stop()
			if stop {
				fmt.Println("Timer stopped")
			}
			return
		}
	}()
}

// wake me up after specified seconds
func wakeMeUp(timeoutSec int) {
	fmt.Println("===Timer Demo===")
	startTimer(1)
	startTimer(5)
	time.Sleep(time.Second * 5)
	fmt.Println("===Timer Demo done===")
}

// ping me periodically
func pingMePeriodically(timeoutSec int) {
	fmt.Println("===Ticker Demo===")
	ticker := time.NewTicker(time.Second * time.Duration(timeoutSec))
	timer := time.After(time.Second * 3)
	go func() {
		for {
			select {
			case t1 := <- timer:
				fmt.Println("Stopping the pings now", t1)
				ticker.Stop()
				return
			case t2 := <- ticker.C:
				fmt.Println("Ping at", t2)
			}
		}
	}()
	time.Sleep(time.Second * 4)
	fmt.Println("===Ticker Demo Done===")
}

func main() {
	wakeMeUp(2)
	pingMePeriodically(1)
}
