package main

import "testing"

func testBanger(t *testing.T) {
	got := banger()
	want := "miniponi!!!!!!!!"

	if got != want {
		t.Errorf("got %v want %v", got, want)

	}
}

func testBangerUni(t *testing.T) {
	got := bangerUni()
	want := "miniponi!!!!!!!!"

	if got != want {
		t.Errorf("got %v want %v", got, want)

	}
}

//[borkab@severest banger]$ go test -v
//testing: warning: no tests to run
//PASS
//ok      learngo/github/learngo/gitto/banger     0.003s
