package main

import (
	"fmt"
	"time"
)

var pln = fmt.Println
var p = fmt.Print

func timePrint(t time.Time) {
	p(t.Year())
	p("-", t.Month())
	p("-", t.Day())
	p(" ", t.Hour())
	p(":", t.Minute())
	p(":", t.Second())
	p(":", t.Nanosecond())
	p(" ", t.Weekday())
	p(" ", t.Location())
	pln()
}

func timeDemo() {
	pln("timeDemo...")
	t1 := time.Now()
	timePrint(t1)

	t2 := time.Date(2012, 06, 14, 9, 15, 01, 123456789, time.UTC)
	timePrint(t2)
	pln("--------------------------")
}

func main() {
	timeDemo()
}
