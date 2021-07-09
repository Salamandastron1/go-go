package main

import "fmt"

func main() {
	fmt.Println(mySum(2, 3, 4, 5))
}

func mySum(xi ...int) int {
	accumulator := 0
	for _, v := range xi {
		accumulator += v
	}
	return accumulator
}
