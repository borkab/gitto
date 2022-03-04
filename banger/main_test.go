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
