package word

import (
	"fmt"
	"gogo/testing_and_benchmarking/ninja_level_13/2/quote"
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
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}
func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}
func ExampleCount() {
	longString := "meow oink moo"
	fmt.Println(Count(longString))
	// Output:
	// 3
}
