package main

import (
	"fmt"
	"runtime"
)

func main() {
	x := 10
	y := 10
	c := gen(x, y)
	for i := 0; i < x*y; i++ {
		fmt.Println(i, <-c)
	}

	fmt.Println("ROUTINES", runtime.NumGoroutine())
}

// gen creates a channel and go routines that
// add 10 numbers to the channel
func gen(numRoutines, numNumbers int) chan int {
	c := make(chan int)
	for i := 0; i < numRoutines; i++ {
		go func() {
			for x := 0; x < numNumbers; x++ {
				c <- x
			}
		}()
		fmt.Println("ROUTINES", runtime.NumGoroutine())
	
	return c
}
