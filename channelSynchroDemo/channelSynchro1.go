package main

import (
	"fmt"
	"time"
)

func newTask(taskNotify chan bool) {
	fmt.Println("New task Received.")
	fmt.Println("Sending new task notification...")
	taskNotify <- true
}

func beginTask(taskNotify chan bool) {
	fmt.Println("Waiting for new task...")
	newTask := <- taskNotify
	if newTask == true {
		fmt.Println("New task notification recieved.")
	}
}


/*
go channels can be used to synchronize execution of functions across goroutines.
Block on channel to wait for some signal/task/notification etc.
Unblock channel by sending signal/task/notification etc
*/
func taskSyncDemo() {
	// create a channel for task notification
	taskNotify := make(chan bool)
	// block till we have a task to perform
	go beginTask(taskNotify)
	// notify the task channel about a new task in the system
	go newTask(taskNotify)
	// allow some time for routines to run
	time.Sleep(time.Second * 1)
}

func main() {
	taskSyncDemo()
}
