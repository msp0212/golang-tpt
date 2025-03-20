package main

import (
	"fmt"
	"time"
	"strconv"
)

type person struct {
	id int
	name string
}


var unbufferedPersonTransport chan person
var bufferedPersonTransport chan person

func sendPerson(p *person, transport chan person) {
	fmt.Println("Sending person to andromeda galaxy:", *p)
	transport <- *p
}

func recvPerson(transport chan person) {
	p := <- transport
	fmt.Println("Received preson in andromeda galaxy:", p)
}

/*
Unbuffered channel communication.
Does not have a buffer to store data.
Communication is synchronous:
send/recv from channel blocks if corresponding recv/send is not done.
*/ 
func runUnbufferedPersonTransport() {
	fmt.Println("runUnbufferedPersonTransport...")
	p := person{1, "person1"}
	// create a channel for communication between goroutines
	unbufferedPersonTransport = make(chan person)
	fmt.Println("Transport person")
	go sendPerson(&p,  unbufferedPersonTransport)
	go recvPerson(unbufferedPersonTransport)
	time.Sleep(time.Second)
	fmt.Println("Transport completed")
}

/*
Buffered channel communication.
Has a buffer to store data.
Communication is asynchronous:
send/recv do not block as long as buffer is not full/empty
*/ 
func runBufferedPersonTransport() {
	fmt.Println("runBufferedPersonTransport...")
	// create a channel for communication between goroutines
	bufferedPersonTransport = make(chan person, 4)
	fmt.Println("Transport 4 persons")
	for i := range 4 {
		p := person{i, "person" + strconv.Itoa(i)}
		sendPerson(&p, bufferedPersonTransport)
	}
	for _ = range 4 {
		recvPerson(bufferedPersonTransport)
	}
	fmt.Println("Transport completed")
}

func main() {
	runUnbufferedPersonTransport()
	runBufferedPersonTransport()
}
