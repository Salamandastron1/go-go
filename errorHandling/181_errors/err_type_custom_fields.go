package main

import (
	"fmt"
	"log"
)

type norgateMathError struct {
	lat  string
	long string
	err  error
}

func (n norgateMathError) Error() string {
	return fmt.Sprintf("a norgate math error occurred: %v %v %v", n.lat, n.long, n.err)
}

func main() {
	_, err := sqrt(-10.42)
	if err != nil {
		log.Fatal(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		nme := fmt.Errorf("norgate math redux: square roort of negative number")
		return 0, norgateMathError{"50.2289", "-99.4526", nme}
	}
	return 42, nil
}
