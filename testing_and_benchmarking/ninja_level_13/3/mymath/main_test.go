package mymath

import (
	"fmt"
	"testing"
)

// Arithmetic operators are based on the most complex number type

func BenchmarkCenteredAvg(b *testing.B) {
	xi := []int{1, 3, 4, 5, 8}
	for i := 0; i < b.N; i++ {
		CenteredAvg(xi)
	}
}

func ExampleCenteredAvg() {
	xi := []int{1, 3, 4, 5, 8, 11, 12}
	fmt.Println(CenteredAvg(xi))
	// Output:
	// 6.2
}

func TestCenteredAverage(t *testing.T) {
	type test struct {
		data []int
		want float64
	}
	cases := []test{
		{
			[]int{1, 1, 1, 1},
			1,
		},
		{
			[]int{2, 6, 7, 8},
			6.5,
		},
		{
			[]int{12, 20, 45, 46, 75, 200},
			46.5,
		},
	}

	for _, v := range cases {
		got := CenteredAvg(v.data)
		if got != v.want {
			t.Errorf("Got: %v\nWant: %v", got, v.want)
		}
	}
}
