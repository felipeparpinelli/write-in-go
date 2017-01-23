package main

import ( 
	"fmt"
	"time"
)

func print(n int) {
    for i := 0; i < 3; i++ {
        fmt.Printf("%d ", n)
        time.Sleep(200 * time.Millisecond)
    }
}

/*
By simply adding the 'go' keyword, 
this new version runs the two calls concurrently and produces a completely different result: 3 2 3 2 3 2.
It is important to note that goroutines depend on the main () function to continue running. 
In other words, goroutines die when the main () function ends.
*/

func main() {
	go print(2)
    print(3)
}