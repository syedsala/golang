package main

import "fmt"

// Function closures is a function value that references variables from outside its body.
// The function may access and assign values to the referenced variables.

func adder() func(int) int {
	sum := 0
	// return a closure function
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos := adder()
	for i := 0; i < 5; i++ {
		fmt.Println(pos(i))
	}
}

/*
0
1
3
6
10
 */