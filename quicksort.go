package main

import (
    "fmt"
	"os"
	"strconv"
)

// To run: "go run quicksort.go 77 121 11 10 3 49 1 2 64 4 22 32 88"
func main() {
    input := os.Args[1:]
    arr := make([]int, len(input))
    for i, n := range input {
        number, err := strconv.Atoi(n)
        if err != nil {
            fmt.Printf("%s not a valid number!\n", n)
			os.Exit(1) 
		}
        arr[i] = number
    }
    fmt.Println(quicksort(arr))
}

// Quick Sort: Runtime: O(n log(n)) average, O(n2) worst case. Memory: O(log(n)). 
func quicksort(numbers []int) []int {
    if len(numbers) <= 1 {
    	return numbers
	}

	// Create a copy of the original list
	n := make([]int, len(numbers))
	copy(n, numbers)

	// Choose the half-list element as pivot
	pivotIndex := len(n) / 2
	pivot := n[pivotIndex]

	// Remove the pivot
	n = append(n[:pivotIndex], n[pivotIndex+1:]...)

	left, right := partition(n, pivot)

	return append(append(quicksort(left), pivot), quicksort(right)...)
}

func partition (numbers []int, pivot int) (left []int, right []int) {

	// Known as 'blank identifier'. In Go, an unused variable
	// causes compilation error. As 'range' returns the element and your index,
	// we need to use 'blank' because the returned index will not be used.
    for _, n := range numbers {
        if n <= pivot {
            left = append(left, n)
        } else {
            right = append(right, n)
        }
	}
    
    return left, right
}