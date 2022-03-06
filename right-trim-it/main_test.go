package main

import (
	"reflect"
	"testing"
)

func TestRighgtTrimIt(t *testing.T) {
	n := "milka            "
	got := rightTrimIt(n)
	want := "milka"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}

func TestSplit(t *testing.T) {
	n := "milka"
	got := Split(n, "")
	want := []string{"m", "i", "l", "k", "a"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)

	}
	/*
		[borkab@severest right-trim-it]$ go test
		--- FAIL: TestSplit (0.00s)
			main_test.go:25: got [m i l k a] want [m i l k a]
		FAIL
		exit status 1
		FAIL    learngo/github/learngo/gitto/right-trim-it      0.004s
	*/
}

func TestLenSplit(t *testing.T) {
	n := []string{"m", "i", "l", "k", "a"}
	got := lenSplit(n)
	want := 5

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

}
