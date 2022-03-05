package main

import (
	"fmt"
	"strings"
)

// EXERCISE: Right Trim It
//
//  1. Look at the documentation of strings package
//  2. Find a function that trims the spaces from
//     only the right-most part of the given string
//  3. Trim it from the right part only
//  4. Print the number of characters it contains.
//
// RESTRICTION
//  Your program should work with unicode string values.

func main() {
	m := "tunyacsap             "
	fmt.Println(lenSplit(Split((rightTrimIt(m)), "")))
	//fmt.Println(Split((rightTrimIt(m)), ""))
}

func rightTrimIt(name string) string {
	s := strings.TrimRight(name, " ")
	return s
}

func Split(s string, sep string) []string {

	sp := strings.Split(s, "")
	return sp
}

func lenSplit(sp []string) int {
	ls := len(sp)
	return ls
}

/*
func main() {
	name := "inanç           "

	name = strings.TrimRight(name, " ")
	l := utf8.RuneCountInString(name)

	fmt.Println(l)
}
*/
