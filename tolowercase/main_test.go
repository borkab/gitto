package main

import "testing"

func TestToLowercase(t *testing.T) {
	input := "TUROK"

	got := ToLowercase(input)
	want := "turok"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
