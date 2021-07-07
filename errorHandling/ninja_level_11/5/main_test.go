package main

import (
	"testing"
)

func TestSimple(t *testing.T) {
	expected := 1
	actual := returnOne()

	if actual != expected {
		t.Error("Not the actual not equaling the expected")
	}
}
