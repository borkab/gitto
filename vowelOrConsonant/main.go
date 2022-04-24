package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	l := len(args[1])
	s := args[1]

	if l != 1 {
		fmt.Println("Give me a letter")
		return
	}

	if s == "a" || s == "e" || s == "i" || s == "o" || s == "u" {
		fmt.Printf("%q is a vowel\n", s)
	} else if s == "w" || s == "y" {
		fmt.Printf("%q is sometimes a vowel, sometimes not\n", s)
	} else {
		fmt.Printf("%q is a consonant\n", s)
	}
}
