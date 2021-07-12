// Package acdc gets you ready to rock
package acdc

// Sum adds any amount of numbers together
func Sum(xi ...int) int {
	accumulator := 0
	for _, v := range xi {
		accumulator += v
	}
	return accumulator
}
