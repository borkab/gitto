package main

import "testing"

func TestTrimIt(t *testing.T) {
	text := `

	Luke, 
I am your father
	`
	got := trimIt(text)
	want := "Luke, \nI am your father"

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

}
