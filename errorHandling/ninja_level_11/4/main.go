package main

import (
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		// can also use errors.New() to generate a message
		// of type error
		msg := fmt.Errorf("negative numbers do not have a square root: %v", f)
		return 0, sqrtError{"40", "20", msg}
	}
	return 42, nil
}
