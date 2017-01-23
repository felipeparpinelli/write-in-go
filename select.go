package main

import (
	"fmt"
)


func odds_or_evens_channels(nums []int, odds, evens chan<- int, ready chan<- bool) {
    for _, n := range nums {
        if n % 2 == 0 {
            evens <- n
		} else {
			odds <- n
		} 
	}
    ready <- true
}

func main () {

	odds, evens := make(chan int), make(chan int)
	ready := make(chan bool)
	nums := []int{1, 23, 42, 5, 8, 6, 7, 4, 99, 100}

	go odds_or_evens_channels(nums, odds, evens, ready)

	var oddList, evenList []int
	end := false

	for !end {
		select {
		case m := <- odds:
			oddList = append(oddList, m)
		case m := <- evens:
			evenList = append(evenList, m)
		case end = <- ready:
		}
	}
	
	fmt.Printf("Odds: %v | Evens: %v\n", oddList, evenList)
}