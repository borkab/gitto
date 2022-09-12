package main

import "testing"

func TestSum(t *testing.T) {
	a := 0
	got := sum(a)
	want := 55

	if got != want {
		t.Errorf("got: %v want %v", got, want)
	}
}
