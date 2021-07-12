package main

import "testing"

func TestMySum(t *testing.T) {
	type test struct {
		data     []int
		expected int
	}
	cases := []test{
		{[]int{1, 2, 3, 4, 5}, 15},
		{[]int{5, 5, 5}, 15},
		{[]int{1, 1, 1, 1, 1}, 5},
	}

	for _, v := range cases {
		got := mySum(v.data...)
		if v.expected != got {
			t.Errorf("Expected: %v Got: %v", v.expected, got)
		}
	}
}
