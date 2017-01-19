package main

import ( 
	"fmt"
	)

func printData(name string, age int) {
	fmt.Printf("name: %s  age: %d", name, age)
}

func sum (n, m int) int {
	return n + m
}

func main () {
	printData("Fulano", 10)
	s := sum(3, 5)
	fmt.Println(" The sum is", s)
}