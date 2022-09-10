package main

import "testing"

func TestSumV(t *testing.T) {
	var x int
	got := testSumV(x)
	want := "1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9 + 10 = 55"

	if got != want {
		t.Error("got %d want %v", got, want)
	}
}
