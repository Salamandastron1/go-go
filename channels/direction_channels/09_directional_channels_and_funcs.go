package main

import "fmt"

func main() {
	c := make(chan int)

	//send
	go send(c)
	//recieve
	recieve(c)

	fmt.Println("about to exit")
}

//send
func send(c chan<- int) {
	c <- 42
}

//recieve
func recieve(c <-chan int) {
	fmt.Println(<-c)
}
