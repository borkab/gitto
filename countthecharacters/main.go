package main

import (
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	text := os.Args[1]
	fmt.Println(countRunes(text))
	fmt.Println(countChars(text))

}

func countChars(text string) int {

	length := len(text)
	return length

}

func countRunes(text string) int {

	lenght := utf8.RuneCountInString(text)
	return lenght
}
