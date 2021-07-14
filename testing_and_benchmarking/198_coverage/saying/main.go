package saying

import "fmt"

func Greet(s string) string {
	return fmt.Sprint("Welcome my dear ", s)
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
