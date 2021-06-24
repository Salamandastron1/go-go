package main

import "fmt"

func main() {
	c := make(chan int)
	cr := make(<-chan int) //recieve <-chan
	cs := make(chan<- int) // sned chan<-

	fmt.Println("---------------")
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cs\t%T\n", cs)

	//specific to general doesn't convert
	fmt.Println("--------")
	fmt.Printf("c\t%T\n", (<-chan int)(c))
	fmt.Printf("c\t%T\n", (chan<- int)(c))

	fmt.Println("SPECIFIC TO GENERAL WORKS, GENERAL TO SPECIFIC DOES NOT")
}
