package main

import "fmt"

func main() {
	quit := make(chan bool)

	go func() {
		for {
			select {
			case <-quit:
				fmt.Println("Received true value in goroutine")
				return
			default:
				fmt.Println("Received false value in goroutine")
			}
		}
	}()
	fmt.Println("Passing true value in channel")
	quit <- true
	fmt.Println("Passed true value in channel")
}
/*
Passing true value in channel
Received false value in goroutine
Received true value in goroutine
Passed true value in channel
*/