// Package mymath provides ACME inc math solutions
// want to name folders the name of the package
package mymath

// Sum adds an unlimited number of values of type int
func Sum(xi ...int) int {
	sum := 0

	for _, v := range xi {
		sum += v
	}
	return sum
}
