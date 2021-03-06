package main

import "fmt"

func main() {
	c := createChannels()

	for k := 0; k < 100; k++ {
		v, ok := <-c
		fmt.Println(v, ok)
	}

	close(c)
	v, ok := <-c
	fmt.Println(v, ok)
}

// createChannels makes 10 new channels
// returns a slice[]chan int
func createChannels() chan int {
	c := make(chan int)

	for i := 0; i < 10; i++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
		}()
	}
	return c
}
