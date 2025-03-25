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
SELECT, NON-BLOCKING CHANNEL OPS, TIMEOUTS & CLOSING CHANNELS:
Select:
makes a goroutine block until one of the case statements can run.
one case statement is chosen at random if multiples are ready to run
Non-Blocking/Default case:
select can be made non-blocking by including a default case.
if none of the case statements are ready to run, default case is executed
Timeouts:
case statement can be made to wait for timeouts using time. 
this is important to enforce time bounds in case of communication with external
resources
Closing a channel:
Indicates that no more payloads will be sent over the channel.
Receive can still be done on the channel.
Once all the payloads are received, indication is given to the receiver about
the closure via a return value
*/

/*
CHANNEL DIRECTION:
Direction can be specified when channels are used as func params.
Adds extra type safety to the function and avoids programmer errors
*/
// non blocking send using select with a default case
func newTask(taskNotify chan <- string, taskNum int) {
	// send direction specified in the func param
	// compile time error is generated if a recv is attempted on this param
	task := "task" + strconv.Itoa(taskNum)
	fmt.Println("New task", task)
	select {
	case taskNotify <- task:
		fmt.Println("Sent new task notification")
	default:
		fmt.Println("Failed sending new task notification")
	}
}

func receiveTask(taskNotify <- chan string) {
	// recv direction specified in the func param
	// compile time error is generated if a send is attempted on this param
	// taskNotify <- false // compile time error
	fmt.Println("Waiting for new task...")
	task := <- taskNotify
	fmt.Println("New task recieved", task)
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
	taskNotify := make(chan string)
	// block till we have a task to perform
	go receiveTask(taskNotify)
	time.Sleep(time.Second * 1)
	// notify the task channel about a new task in the system
	go newTask(taskNotify, 0)
	// allow some time for routines to run
	time.Sleep(time.Second * 1)
	fmt.Println("===Task Synchronisation Done===")
}


// send new tasks with specified frequency upto maxTasks
func newTaskSender(taskNotify chan <- string, freqSec int, maxTasks int) {
	for i := 0; i < maxTasks; i++ {
		newTask(taskNotify, i)
		time.Sleep(time.Second * time.Duration(freqSec))
	}
	close(taskNotify)
}

// blocking receive with a specified time out
func runTaskLoop(taskNotify <- chan string, timeoutSec int) {
TaskLoop:
	for {
		fmt.Println("Run Task Loop")
		select {
		case task, open := <- taskNotify:
			fmt.Println("Received new task", task)
			if open == false {
				fmt.Println("Task Notify channel closed")
				break TaskLoop
			}
		case <- time.After(time.Second * time.Duration(timeoutSec)):
			fmt.Println("Timed out waiting for new tasks")
			break TaskLoop
		}
	}
	fmt.Println("Breaking out from task loop")
}

func taskLoopDemo() {
	fmt.Println("===Task System Start===")
	taskNotify := make(chan string, 1)
	go runTaskLoop(taskNotify, 3)
	go newTaskSender(taskNotify, 2, 1)
	// let the task system run for 10 seconds
	time.Sleep(time.Second * 10)
	fmt.Println("===Task System Stop===")
}

/*
Range over channels:
range is supported over channels also.
range iteration terminates when channel is closed
*/
func channelRangeDemo() {
	fmt.Println("===Range over channel===")
	c := make(chan string, 5)
	c <- "str1"
	c <- "str2"
	c <- "str3"
	close(c)

	for s := range c {
		fmt.Println(s)
	}
	fmt.Println("===Range over channel done===")
}


func main() {
	runUnbufferedPersonTransport()
	runBufferedPersonTransport()
	taskSyncDemo()
	taskLoopDemo()
	channelRangeDemo()
}
