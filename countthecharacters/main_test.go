package main

import (
	"testing"
)

func TestCountChars(t *testing.T) {
	text := "Borka"
	got := countChars(text)
	want := 5

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestCountRunes(t *testing.T) {
	text := "szeretetunikornis"
	got := countRunes(text)
	want := 17

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}

}

/*
[borkab@severest countthecharacters]$ go test
# learngo/github/learngo/gitto/countthecharacters [learngo/github/learngo/gitto/countthecharacters.test]
./main_test.go:9:19: not enough arguments in call to countChars
        have ()
        want (string)
./main_test.go:18:19: not enough arguments in call to countRunes
        have ()
        want (string)
FAIL    learngo/github/learngo/gitto/countthecharacters [build failed]
*/
