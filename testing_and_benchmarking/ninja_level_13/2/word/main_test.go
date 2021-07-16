package word

import (
	"fmt"
	"testing"
)

func TestUseCount(t *testing.T) {
	s := "Lookey here mister mister"
	want := map[string]int{"Lookey": 1, "here": 1, "mister": 2}
	got := UseCount(s)
	for v := range want {
		if want[v] != got[v] {
			t.Errorf("Got: %v\nWant: %v", got, want)
		}
	}
}
func TestCount(t *testing.T) {
	s := "Oink meow moo, womp"
	want := 4
	got := Count(s)
	if want != got {
		t.Errorf("Got: %v\nWant: %v", got, want)
	}
}
func BenchmarkCount(b *testing.B) {
	longString := "Wow wow, look meow"
	for i := 0; i < b.N; i++ {
		Count(longString)
	}
}
func BenchmarkUseCount(b *testing.B) {
	longString := "Wow wow, look meow"
	for i := 0; i < b.N; i++ {
		UseCount(longString)
	}
}
func ExampleCount() {
	longString := "meow oink moo"
	fmt.Println(Count(longString))
	// Output:
	// 3
}
