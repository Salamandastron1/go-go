package main

import (
	"testing"
)

func TestSimple(t *testing.T) {
	expected := 1
	actual := returnOne()

	if actual != expected {
		t.Errorf("Not the actual: %v not equaling the expected: %v", actual, expected)
	}
}
