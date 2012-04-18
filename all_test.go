package main

import (
	"testing"
)

func TestFib(t *testing.T) {
	correct := []int64{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	for i, v := range correct {
		ans := Fib(i+1)
		if ans != v {
			t.Errorf("Number %d returned by Fib was not %d, but %d", i+1, v, ans)
		}
	}
}
