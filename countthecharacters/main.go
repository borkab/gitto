package main

import (
	"fmt"
	"os"
)

func main() {
	// Currently it returns 7
	// Because, it counts the bytes...
	// It should count the runes (codepoints) instead.
	//
	// When you run it with "borka", it should return 5 not 7.
	length := len(os.Args[1])
	fmt.Println(length)
}
