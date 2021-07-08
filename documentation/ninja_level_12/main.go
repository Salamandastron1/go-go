package main

import (
	"fmt"
	"gogo/documentation/ninja_level_12/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	doge := canine{
		"beefer",
		dog.Years(11),
	}
	fmt.Println(doge)
}
