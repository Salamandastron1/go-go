package main

import "fmt"

// func literal
func main() {
	c := make(chan int)
	go func() {
		c <- 42
	}()
	fmt.Println("I am an anonymous function")
	fmt.Println(<-c)

	bufferedChannel()
}

// buffered channel
func bufferedChannel() {
	c := make(chan int, 1)

	c <- 42

	fmt.Println("I am a buffered channel")
	fmt.Println(<-c)
}
