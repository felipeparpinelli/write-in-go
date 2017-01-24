package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

/*
Previously, we saw an example of how to expect a goroutine to perform its execution through the use of a channel. 
However, when we need to wait for the execution of multiple goroutines, 
manually checking how many of them have ended may become a failed task.
Go provides a type called WaitGroup, present in the sync package, which makes this task much simpler.
*/

func main() {
    timestamp := time.Now()
    rand.Seed(timestamp.UnixNano())
    var control sync.WaitGroup
    for i := 0; i < 5; i++ {
        control.Add(1)
        go execute(&control)
    }
    control.Wait()
    fmt.Printf("stopped with %s.\n", time.Since(timestamp))
}
	
func execute(control *sync.WaitGroup) {
    defer control.Done()
    duration := time.Duration(1+rand.Intn(5)) * time.Second
    fmt.Printf("sleeping for %s...\n", duration)
    time.Sleep(duration)
}