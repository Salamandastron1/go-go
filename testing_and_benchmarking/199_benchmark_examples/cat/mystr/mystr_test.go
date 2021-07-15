package mystr

import (
	"fmt"
	"strings"
	"testing"
)

func TestJoin(t *testing.T) {
	want := "I like meows"
	got := Join([]string{"I", "like", "meows"})
	if got != want {
		t.Error("Got", got, "does not equal want", want)
	}
}
func TestCat(t *testing.T) {
	want := "I like meows"
	got := Cat([]string{"I", "like", "meows"})
	if got != want {
		t.Error("Got", got, "does not equal want", want)
	}
}
func ExampleCat() {
	s := "Shaken not stirred"
	xs := strings.Split(s, " ")
	fmt.Println(Cat(xs))
	// Output:
	// Shaken not stirred
}
func ExampleJoin() {
	s := "Shaken not stirred"
	xs := strings.Split(s, " ")
	fmt.Println(Join(xs))
	// Output:
	// Shaken not stirred
}

const s = "To be or not to be that is the question"

var xs []string

func BenchmarkCat(b *testing.B) {
	xs = strings.Split(s, " ")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Cat(xs)
	}
}
func BenchmarkJoin(b *testing.B) {
	xs = strings.Split(s, " ")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Join(xs)
	}
}
