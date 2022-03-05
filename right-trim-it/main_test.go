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

func TestSplit(t *testing.T) {
	n := "milka"
	got := Split(n, "")
	want := []string{m, i, l, k, a} // [m i l k a] hogy irjam be h ezt elfogadja?

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestLenSplit(t *testing.T) {
	n := []string{m, i, l, k, a} //hogy irjam be a slice-t?
	got := lenSplit(n)
	want := 5

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

}