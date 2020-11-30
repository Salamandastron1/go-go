package main

// channels block
// channels block
// channels blockx
import "fmt"

func main() {
	c := make(chan int, 1)
	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}
