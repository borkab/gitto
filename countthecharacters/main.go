package main

import (
	"os"
	"unicode/utf8"
)

func main() {
	countRunes()
	countChars()
}

func countChars() int {
	// Currently it returns 7
	// Because, it counts the bytes...
	// It should count the runes (codepoints) instead.
	//
	// When you run it with "borka", it should return 5 not 7.
	length := len(os.Args[1])
	return length

	//but it doesn't return 7 as written in the exercise
	//the output is correct, 5
}

func countRunes() int {

	lenght := utf8.RuneCountInString(os.Args[1])
	return lenght
}
