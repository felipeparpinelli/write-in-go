package main

import (
	"fmt"
)

/*
The ability to run different goroutines concurrently is very important 
and opens up several possibilities for solving problems that require high performance 
in the use of processing resources.
However, it is very rare that goroutines are triggered independently, 
without any communication between them. 
The channels were created as an abstraction to enable this communication.
A channel conducts information of a certain type - any valid type in Go.
*/

func main () {
	
	// Defining a channel
	c := make(chan int)

	/*
	To interact with a channel, we use the <- operator (known as the arrow operator). 
	The position of the arrow indicates the direction of communication flow. 
	For example, to send int values to channel c, we use the following notation:
	*/
	
	// c <- 33

	// To receive a value sent by channel c:
	
	// valor := <-c

	// sample:

	go generate(c)

	valor := <-c

	fmt.Println(valor)

}

func generate(c chan int) {
    c <- 33
}

/* 

Initially we created a channel to traffic values of type int.
We then fire a goroutine by executing the generate () function, 
which receives a channel as an argument and simply sends an integer to the received channel.
By default, send and receive operations on one channel block until the other side is ready. 
This fact allows the communication between two goroutines to guarantee the synchronization between them, 
without any locking mechanism being necessary.
For this reason, the next line of the main () function - which receives a channel value - 
will cause the main execution line to be locked until some value is sent to channel c. 
As soon as the value 33 is sent by the generate () function, the main execution line will then be unlocked, 
the value 33 will be consumed, assigned to the variable 'value' and printed on the console.

*/


// Bidirectional channel
func bidirectional(c chan int) {
	c <- 1
    c <- 2 
    c <- 3
	
	close(c)
}

// We can define a channel with only one direction
func oneDirection(c chan <- int) {
	c <- 1
	
	close(c)
}

// And we can define a read-only channel
func readOnly(c <- chan int) {
	fmt.Println(c)
}