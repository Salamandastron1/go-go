package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	want := 77
	got := Years(11)
	if want != got {
		t.Errorf("Got: %v\nWant: %v", got, want)
	}
}
func TestYearsTwo(t *testing.T) {
	want := 77
	got := YearsTwo(11)
	if want != got {
		t.Errorf("Got: %v\nWant: %v", got, want)
	}
}

func ExampleYears() {
	fmt.Println(Years(11))
	// Output:
	// 77
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(11))
	// Output:
	// 77
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(11)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(11)
	}
}
