package main

import (
	"fmt"
	"File"
)

// Defining a Interface
type Operation interface {
    Calculate() int
}

func main() {

	// Declaring
	file := File{size: 1.12, name: "pgr.go", chars: 12986, words: 1862, lines: 220}
	fmt.Println(file)

	// Accessing the values
	fmt.Printf("%s\t%.2fKB\n", file.name, file.size)

	// Updating a value
	file.name = "pgr_modified.go"
	fmt.Printf("%s\t%.2fKB\n", file.name, file.size)

	// Call Struct methods
	fmt.Printf("MeanWordsPerLine: %.2f\n", file.MeanWordsPerLine())
    fmt.Printf("MeanWordSize: %.2f\n", file.MeanWordSize())	

}