package main

import (
	"fmt"
	"log"
)

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Fatal(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		// format variables for errors using the fmt package
		return 0, fmt.Errorf("norgate math: square root of negative number: %v", f)
	}
	return 42, nil
}
