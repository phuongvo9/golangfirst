package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	got := Sum([5]int{1, 2, 3, 4, 5})
	want := 15
	if got != want {
		t.Errorf("got %d want %d have %v", got, want, [5]int{1, 2, 3, 4, 5})
	}
}
