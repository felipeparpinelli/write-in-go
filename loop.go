package main
import (
    "fmt"
	"math/rand"
	"time"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    n := 0
	for 
	{ 
		n++
        i := rand.Intn(4200)
        fmt.Println(i)
		if i%42 == 0 {
			break
		} 
	}
    fmt.Printf("%d interactions.\n", n)

    var i int
    firstLoop: //Naming the block
	for {
    	for i = 0; i < 10; i++ {
        	if i == 5 {
            	break firstLoop
			}
		}
	}
	fmt.Println(i)

	var j int
	loop: //Name for the loop block
	for j = 0; j < 10; j++ {
    	fmt.Printf("for i = %d\n", j)
    	switch j {
    		case 2, 3:
        		fmt.Printf("break switch, j == %d.\n", j)
				break 
			case 5:
        		fmt.Println("break loop, j == 5.")
				break loop 
		}
	}
	fmt.Println("end")
}
