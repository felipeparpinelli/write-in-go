package main
import "fmt"

// Defining Struct
type File struct {
    name	string
    size	float64
    chars	int
    words	int
    lines	int
}

/*
We can also define methods whose receivers are structs, 
achieved in practice an effect very similar to that of a class.
example:
*/
func (file *File) MeanWordSize() float64 {
    return float64(file.chars) / float64(file.words)
}

func (file *File) MeanWordsPerLine() float64 {
	return float64(file.words) / float64(file.lines)
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