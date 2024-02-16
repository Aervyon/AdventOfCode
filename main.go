package main

import (
	"fmt"

	"github.com/Aervyon/AdventOfCode/day1"
)

func main() {
	fmt.Println("Please enter the problem you'd like to do. Available options: 1-50")

	var problem int
	_, err := fmt.Scanln(&problem)
	if err != nil {
		fmt.Printf("Error reading your input, %q\n", err)
		return
	}

	switch problem {
	case 1:
		day1.Trebuchet()
	case 2:
		day1.Trebuchet2()
	default:
		println("Sorry we haven't made that yet")
	}
}
