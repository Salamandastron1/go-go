package mymath2

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	type test struct {
		data     []int
		expected int
	}
	cases := []test{
		{[]int{21, 21}, 42},
		{[]int{1, 2, 3}, 6},
		{[]int{-1, 0, -2}, -3},
		{[]int{-1, 0, 1}, 0},
	}
	for _, v := range cases {
		actual := Sum(v.data...)
		if actual != v.expected {
			t.Errorf("Actual: %v\n does not equal Expected: %v", actual, v.expected)
		}
	}
}

func ExampleSum() {
	fmt.Println(Sum(1, 3, -3, 2))
	// Output:
	// 3
}
