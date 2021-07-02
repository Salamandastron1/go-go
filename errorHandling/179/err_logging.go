package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer foo()
	_, err := os.Open("no-file.txt")
	if err != nil {
		// fmt provides basic text output no frills
		// fmt.Println("err happened", err)
		// log prints date timestamp
		// log.Println("err happened", err)
		// fatal calls os.Exit(1) when called
		log.Fatal(err)
	}
}

func foo() {
	fmt.Println("When os.Exit() is called deferred functions don't run")
}
