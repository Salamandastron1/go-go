package main

import "fmt"

func main() {
	c := make(chan int)
	cr := make(<-chan int) //recieve
	cs := make(chan<- int) //send

	c <- 42
	// cr <- 43

	fmt.Println("------")
	fmt.Printf("c\t%T\n%v\n", c, <-c)
	fmt.Printf("cr\t%T\n%v\n", cr, cr)
	fmt.Printf("cs\t%T\n", cs)
}
