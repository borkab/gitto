package main

import "testing"

func TestBanger(t *testing.T) {
	msg := "Luke"
	got := banger(msg)
	want := "Luke!!!!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestBangerUni(t *testing.T) {
	msg := "Skywalker"
	got := bangerUni(msg)
	want := "Skywalker!!!!!!!!!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

/*
[borkab@severest banger]$ go test -v
# learngo/github/learngo/gitto/banger [learngo/github/learngo/gitto/banger.test]
./main_test.go:7:15: too many arguments in call to banger
        have (string)
        want ()
./main_test.go:18:18: too many arguments in call to bangerUni
        have (string)
        want ()
FAIL    learngo/github/learngo/gitto/banger [build failed]
*/
