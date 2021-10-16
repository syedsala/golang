package main

import "fmt"

/*
String literals raw( \ characters don't have special meaning) & interpreted
 */
func main() {
	var (
		raw = `This\this\ta\traw\tstring`
		interpreted = "This\tis\ta\tinterpreted\tstring"
	)
	fmt.Println(raw)
	fmt.Println(interpreted)
}
