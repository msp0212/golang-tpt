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


func sendPerson(p *person, transport chan person) {
	fmt.Println("Sending person to andromeda galaxy:", *p)
	transport <- *p
}

func recvPerson(transport chan person) {
	p := <- transport
	fmt.Println("Received preson in andromeda galaxy:", p)
}

/*
UNBUFFERED CHANNEL COMMUNICATION:
Does not have a buffer to store data.
Communication is synchronous:
send/recv from channel blocks if corresponding recv/send is not done.
*/ 
func runUnbufferedPersonTransport() {
	fmt.Println("===runUnbufferedPersonTransport===")
	p := person{1, "person1"}
	// create a channel for communication between goroutines
	unbufferedPersonTransport := make(chan person)
	fmt.Println("Transport person")
	go sendPerson(&p,  unbufferedPersonTransport)
	go recvPerson(unbufferedPersonTransport)
	time.Sleep(time.Second)
	fmt.Println("===Transport completed===")
}

/*
BUFFERED CHANNEL COMMUNICATION:
Has a buffer to store data.
Communication is asynchronous:
send/recv do not block as long as buffer is not full/empty
*/ 
func runBufferedPersonTransport() {
	fmt.Println("===runBufferedPersonTransport===")
	// create a channel for communication between goroutines
	bufferedPersonTransport := make(chan person, 4)
	fmt.Println("Transport 4 persons")
	for i := range 4 {
		p := person{i, "person" + strconv.Itoa(i)}
		sendPerson(&p, bufferedPersonTransport)
	}
	for _ = range 4 {
		recvPerson(bufferedPersonTransport)
	}
	fmt.Println("===Transport completed===")
}

/*
CHANNEL DIRECTION:
Direction can be specified when channels are used as func params.
Adds extra type safety to the function and avoids programmer errors
*/
func newTask(taskNotify chan <- bool) {
	// send direction specified in the func param
	// compile time error is generated if a recv is attempted on this param
	fmt.Println("New task Received.")
	fmt.Println("Sending new task notification...")
	taskNotify <- true
}

func beginTask(taskNotify <- chan bool) {
	// recv direction specified in the func param
	// compile time error is generated if a send is attempted on this param
	// taskNotify <- false // compile time error
	fmt.Println("Waiting for new task...")
	newTask := <- taskNotify
	if newTask == true {
		fmt.Println("New task notification recieved.")
	}
}


/*
CHANNEL SYNCHRONISATION:
go channels can be used to synchronize execution of functions across goroutines.
Block on channel to wait for some signal/task/notification etc.
Unblock channel by sending signal/task/notification etc
*/
func taskSyncDemo() {
	fmt.Println("===Task Synchronisation Demo===")
	// create a channel for task notification
	taskNotify := make(chan bool)
	// block till we have a task to perform
	go beginTask(taskNotify)
	// notify the task channel about a new task in the system
	go newTask(taskNotify)
	// allow some time for routines to run
	time.Sleep(time.Second * 1)
	fmt.Println("===Task Synchronisation Done===")
}

func main() {
	runUnbufferedPersonTransport()
	runBufferedPersonTransport()
	taskSyncDemo()
}
