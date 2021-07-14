package saying

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	got := Greet("James")
	want := "Welcome my dear James"

	if got != want {
		t.Error("got", got, "want", want)
	}
}

func ExampleGreet() {
	fmt.Println(Greet("Tim"))
	// Output:
	// Welcome my dear Tim
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("James")
	}
}
