package main

import (
    "fmt"
	"regexp"
	"strings"
)

func main() {
    expr := regexp.MustCompile("\\b\\w")
    transform := func(s string) string {
        return strings.ToUpper(s)
	}

    text := "antonio carlos jobim"
    fmt.Println(transform(text))

    fmt.Println(expr.ReplaceAllStringFunc(text, transform))

    // OR

    fmt.Println(expr.ReplaceAllStringFunc(text, func(s string) string {
        return strings.ToUpper(s)
	}))
}
