package main

import ( 
	"fmt"
	"os"
	)

func CreateFiles(dirBase string, files ...string) {
    for _, name := range files {
        filePath := fmt.Sprintf(
            "%s/%s.%s", dirBase, name, "txt")
        file, err := os.Create(filePath)
        defer file.Close()
        if err != nil {
            fmt.Printf("Error creating file %s: %v\n",
                name, err)
            os.Exit(1)
		}
        fmt.Printf("Created file.\n", file.Name())
    }
}


func main() {
    tmp := os.TempDir()
    CreateFiles(tmp)
    CreateFiles(tmp, "test1")
    CreateFiles(tmp, "test2", "test3", "test4")
}