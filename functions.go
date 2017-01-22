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


func PrecoFinal(precoCusto float64) (precoDolar float64, precoReal float64) {
    fatorLucro := 1.33
    taxaConversao := 2.34
    precoDolar = precoCusto * fatorLucro
    precoReal = precoDolar * taxaConversao
	return
}

func main () {
	printData("Fulano", 10)
	s := sum(3, 5)
	fmt.Println(" The sum is", s)
}