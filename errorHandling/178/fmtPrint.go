package main

import "fmt"

func main() {
	// Println returns number of bytes written
	// and an error
	n, err := fmt.Println("Hello")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)

}
