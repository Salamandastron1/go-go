package mystr

import "strings"

// Cat prints out a string
func Cat(xs []string) string {
	s := xs[0]
	for _, v := range xs[1:] {
		s += " "
		s += v
	}
	return s
}

// Join joins a array of strings together based on delimeter
func Join(xs []string) string {
	return strings.Join(xs, " ")
}
