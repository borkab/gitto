package main

import (
	"testing"
)

func testMain(t *testing.T) {

	got := countChars()
	want := 5 //input:borka

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
func testCountRunes(t *testing.T) {

	got := countRunes()
	want := 5 //input:borka
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}

}

//output: pass ok
//but warning: no tests to run
