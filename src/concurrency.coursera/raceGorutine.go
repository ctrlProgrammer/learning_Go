package main

import (
	"time"
)

var racedVariable int = 0
var counter int = 0

// Note: in this function the program set a new value on the variable
// In will be executed 100 times so it will find 50 even numbers
func writer() {
	for i := 0; i < 100; i++ {
		racedVariable += 1
	}
}

// Note: In this function the program read the racedVariable, for each even it will be reseted and the counter will increase
// It means that the program read and write in the same global scope variable
func reader() {
	for i := 0; i < 100; i++ {
		var getter int = racedVariable

		if getter%2 == 0 {
			counter += 1
			getter = 0
		}
	}
}

func main() {
	// Not, the explanation about this. It can be a problem because we cant know when the program reads or write the variable so it is possible to seee some jumped values
	// on different executions, for example when the program write the number 2 and then the 3 without reading it, it will jump an even number
	// It is possible because the goroutines not have a deterministic method, so it can run on different time frames while they are running concurrently
	go writer()
	go reader()
	time.Sleep(time.Second * 10)
}
