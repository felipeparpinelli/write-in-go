package main

import (
	"fmt"
	"time"
)

func execute(c chan<- bool) {
	// force timeout: it's expected that the execution takes 2 seconds (see line 26)
    time.Sleep(5 * time.Second)
    c <- true
}

func main() {
	
	c := make(chan bool, 1)
	go execute(c)

	fmt.Println("Waiting...")

	end := false
	for !end {
	    select {
	    case end = <-c:
	        fmt.Println("end!")
	    case <-time.After(2 * time.Second):
	        fmt.Println("Timeout!")
			end = true 
		}
	}
}