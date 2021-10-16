package main

import (
	"fmt"
	"reflect"
)
// golang is typed language, its has strict type requirement
// supports explicit type conversion

func main() {
	a := 5
	b := 5.5
	c := a + int(b)  // get an error mismatched types int and float64
	fmt.Println(c)   // 10

	d := float64(a) + b
	fmt.Println(d)   // 10.5
	fmt.Println(reflect.TypeOf(d)) // float64
}
