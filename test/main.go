package main

import (
	"fmt"
)

// Calculate returns x + 2.
func Calculate(x int) (result int) {
	result = x + 2
	fmt.Println("The result is: ", result)
	return result
}

func main() {
	fmt.Println("Hello World")
	Calculate(10)
}
