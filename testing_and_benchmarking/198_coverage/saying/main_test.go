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

// tools for looking at tests
// -bench looks at the operations and memory consumption
// of designated operations
// go test -bench ./...
// -cover tells how much of your coding is checked by tests
// go test -cover ./...
// -coverprofile outputs coverage report to a specified file
// go test -coverprofile c.out ./...
// command below allows you to view coverage file in html
// in the browser, highlighting covered and uncovered areas
// go tool cover -html=c.out
