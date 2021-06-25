package main

import "fmt"

func main() {
	c := make(chan int)

	//send
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
	//recieve
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

//send
func send(c chan<- int) {

}

//recieve
func recieve(c <-chan int) {
}
