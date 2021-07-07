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
	fmt.Println(c, c.(customErr).name)
}
