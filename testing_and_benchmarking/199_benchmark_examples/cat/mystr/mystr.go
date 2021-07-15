package mystr

import "strings"

// Cat prints out a string
func Cat(xs []string) string {
	newStr := ""
	for i, v := range xs {
		newStr += v
		if i == len(xs)-1 {
			return newStr
		}
		newStr += " "
	}
	return newStr
}

// Join joins a array of strings together based on delimeter
func Join(xs []string) string {
	return strings.Join(xs, " ")
}
