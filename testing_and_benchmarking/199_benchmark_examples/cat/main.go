package main

import (
	"fmt"
	"gogo/testing_and_benchmarking/199_benchmark_examples/cat/mystr"
	"strings"
)

const s = "To be or not to be that is the question"

var xs []string

func main() {
	xs = strings.Split(s, " ")

	for _, v := range xs {
		fmt.Println(v)
	}
	fmt.Printf("\n%s\n", mystr.Cat(xs))
	fmt.Printf("\n%s\n", mystr.Join(xs))
}
