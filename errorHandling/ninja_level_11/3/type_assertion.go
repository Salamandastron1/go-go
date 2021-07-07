package main

import "fmt"

type customErr struct {
	name string
	err  error
}

func (c customErr) Error() string {
	return fmt.Sprintf("%s is to blame for this error: %v", c.name, c.err)
}
func main() {
	cErr := customErr{"tim", nil}
	foo(cErr)
}

func foo(c error) {
	// below at c.(customErr).name is a static type assertion
	// different than conversion, which converts a type
	// assertion allows for specifying which type
	// you know will be passed to the expression
	fmt.Println(c, c.(customErr).name)
}
