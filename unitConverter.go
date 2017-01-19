package main

import ( 
	"fmt"
	"os"
	"strconv"
	)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Use: converter <value> <unit>")
		os.Exit(1)
	}

	unit := os.Args[len(os.Args)-1]
	values := os.Args[1 : len(os.Args)-1]
	var targetUnit string

	if unit == "celsius" {
		targetUnit = "fahrenheit"
	} else if unit == "kilometers" {
		targetUnit = "miles"
	} else {
		fmt.Printf("%s Unknown unit", targetUnit)
		os.Exit(1)
	}

	for i, v := range values {
		value, err := strconv.ParseFloat(v, 64)
		if err != nil {
			fmt.Print("The value %s in position %d it's not a valid number!\n", v, i)
			os.Exit(1)
		}

		var targetValue float64

		if unit == "celsius" {
			targetValue = value * 1.8 + 32
		} else {
			targetValue = value / 1.60934
		}

		fmt.Printf("%.2f %s = %.2f %s\n", value, unit, targetValue, targetUnit)
	}
}