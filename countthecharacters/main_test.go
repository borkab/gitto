package main

import (
	"testing"
)

func TestCountChars(t *testing.T) {

	got := countChars()
	want := 5 //input:borka

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
func TestCountRunes(t *testing.T) {

	got := countRunes()
	want := 5 //input:borka
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}

}

/*--- FAIL: TestMain (0.00s)
    main_test.go:13: got 18 want 5
--- FAIL: TestCountRunes (0.00s)
    main_test.go:21: got 18 want 5
FAIL
exit status 1
FAIL    learngo/github/learngo/gitto/countthecharacters 0.003s
*/
