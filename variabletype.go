package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Double %d is %d\n", v, v*2)
	case string:
		fmt.Printf("%s is %d bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}

/*
Double 21 is 42
hello is 5 bytes long
I don't know type bool!
 */