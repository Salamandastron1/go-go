package main

import "testing"

func TestMySum(t *testing.T) {
	expected := 5
	actual := mySum(2, 3)
	if expected != actual {
		t.Errorf("Expected: %v Got: %v", expected, actual)
	}
}
