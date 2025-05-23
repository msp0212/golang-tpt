package main

import (
	"fmt"
	"errors"
)

/*https://go.dev/blog/go1.13-errors
error is a type in go
errors.New creates a new error obj with an error msg
error like other types can be returned from the functions
nil value denotes no error
note:by convention, errors should be the last return values
*/
func openTheLock(key int) (bool, error) {
	if key < 0 {
		return false, errors.New("Invalid Key")
	} else if key == 2002 {
		return true, nil
	} else {
		return false, errors.New("This key cannot open the lock. Try again!")
	}
}

func openTheDoor() {
	fmt.Println("Open the door...")
	isOpened, err := openTheLock(-1)
	fmt.Println("Lock opened?", isOpened, "error = ", err)
	isOpened, err = openTheLock(101)
	fmt.Println("Lock opened?", isOpened, "error = ", err)
	isOpened, err = openTheLock(2002)
	fmt.Println("Lock opened?", isOpened, "error = ", err)
}


/*
errors can be predeclared
mark certain error conditions in the system
*/
var ErrOutOfMoney = fmt.Errorf("sufficient funds not available")
var ErrTooTired = fmt.Errorf("too tired to travel more")

func letsGoForTravel(funds int, hoursSpentTravelling int) error {
	if funds < 5000 {
		return ErrOutOfMoney
	} else if hoursSpentTravelling > 120 {
		return ErrTooTired
	} else {
		return nil
	}
}

func travel() {
	fmt.Println("Feel like travelling...")
	err := letsGoForTravel(2000, 12)
	if err == nil {
		fmt.Println("Good to go!")
	} else {
		fmt.Println("Can't go:", err)
	}
	err = letsGoForTravel(6000, 150)
	if err == nil {
		fmt.Println("Good to go!")
	} else {
		fmt.Println("Can't go:", err)
	}
	err = letsGoForTravel(6000, 12)
	if err == nil {
		fmt.Println("Good to go!")
	} else {
		fmt.Println("Can't go:", err)
	}
}


/*
custom errors can be created
by implementing Error() method of error interface
note: custom error types usually have "Error" as suffix
*/

type networkError struct {
	fd int
	op int
	failStr string
}

func (ne *networkError) Error() string {
	return fmt.Sprintf("operation = %d on fd = %d failed: %s", ne.op, ne.fd, ne.failStr)
}

func doNetworkOp(fd int, op int) error {
	if fd < 0 {
		return &networkError{fd, op, "invalid file descriptor"}
	} else {
		// network operation done
		return nil
	}
}

func networkOperations() {
	fmt.Println("Network programming is cool...")
	err := doNetworkOp(-1, 2)
	if err != nil {
		fmt.Println("Error while doing network operation: ", err);
	} else {
		fmt.Println("Network operation done successfully")
	}
	err = doNetworkOp(1, 2)
	if err != nil {
		fmt.Println("Error while doing network operation: ", err);
	} else {
		fmt.Println("Network operation done successfully")
	}

}

func main() {
	openTheDoor()
	travel()
	networkOperations()
}
