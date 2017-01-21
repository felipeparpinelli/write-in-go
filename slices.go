package main
import(
	"fmt"
)

//Slices and array
func main () {
	
	//Slices
	fib := []int{1, 1, 2, 3, 5, 8, 13}
	fmt.Println(fib)
	fmt.Println(fib[:3])
	fmt.Println(fib[2:])
	fmt.Println(fib[:])

	/*
	When a slice is created, an array is allocated internally. 
	When we slice this slice and assign it to a new variable, we have a new slice that shares the same internal array as the original. 
	This means that when an element common to both slices is modified, this modification is reflected in the other.
	*/
	originalSlice := []int{1, 2, 3, 4, 5}
	fmt.Println("Original:", originalSlice)
	newSlice := originalSlice[1:3]
	fmt.Println("New slice:", newSlice)
	originalSlice[2] = 13
	fmt.Println("Original slice after modification:", originalSlice)
	fmt.Println("New slice after modification:", newSlice)

	// Other example:
	a := []string{"John", "eats", "at", "home", "daily"}
	b := a[:len(a)-1]
	c := b[:len(b)-1]
	d := c[:len(c)-1]
	e := d[:len(d)-1]
	e[0] = "Felipe"
	fmt.Printf("%v\n%v\n%v\n%v\n%v\n", a, b, c, d, e)

	/* 
	All operations presented above are also valid in an array, 
	and it's important to mention that slicing an array always results in one slice, never another array.
	*/

	// Inserting at the beginning. It's necessary to create another slice.
	s := []int{23, 24, 25}
	n := []int{22}
	s = append(n, s...)
	fmt.Println(s)

	// or, a more elegant solution
	j := []int{23, 24, 25}
	j = append([]int{22}, j...)
	fmt.Println(j)

	// Inserting at the end
	k := make([]int, 0)
	k = append(k, 23)
	fmt.Println(k)

	// Insert at the middle
	l := []int{11, 12, 13, 16, 17, 18}
	v := []int{14, 15}
	l = append(l[:3], append(v, l[3:]...)...)

	// Removing a value at the beginning
	r := []int{20, 30, 40, 50, 60}
	r = r[1:]
	fmt.Println(s)

	// Removing at the end
	t := []int{20, 30, 40, 50, 60}
	t = t[:3]
	fmt.Println(t)

	// Removing at the middle, we should use append again
	u := []int{10, 20, 30, 40, 50, 60}
	u = append(u[:2], u[4:]...)
	fmt.Println(u)

	// Copying a slice, using the 'Copy' method
	numbers := []int{1, 2, 3, 4, 5}
	double := make([]int, len(numbers))
	copy(double, numbers)
	for i := range double {
    	double[i] *= 2
	}
	fmt.Println(numbers)
	fmt.Println(double)

}