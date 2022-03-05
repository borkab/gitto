package main

import "testing"

func TestRighgtTrimIt(t *testing.T) {
	n := "milka            "
	got := rightTrimIt(n)
	want := "milka"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
