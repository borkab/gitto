package main

import (
	"fmt"
	"os"
	"strings"
)

//  1. Look at the documentation of strings package
//  2. Find a function that changes the letters into lowercase
//  3. Get a value from the command-line
//  4. Print the given value in lowercase letters
func main() {
	input := os.Args[1]
	fmt.Println(ToLowercase(input))
}

func ToLowercase(input string) string {
	t := strings.ToLower(input)
	return t
}
