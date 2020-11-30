package main

import "fmt"

func main() {
	c := make(chan int)
	cr := make(<-chan int) //recieve
	cs := make(chan<- int) //send

	fmt.Println("------")
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cs\t%T\n", cs)

	// general to specific assigns
	cr = c
	cs = c

	//specific to general does not
	// c = cr
	// c = cs
	//^^WRONGO
}
