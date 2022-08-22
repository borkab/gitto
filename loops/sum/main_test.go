package main

import "testing"

func TestSum(t *testing.T) {
	var y int
	got := sum(y)
	want := 55

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
