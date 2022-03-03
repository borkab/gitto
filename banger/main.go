package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	//change the Banger program to work with unicode characters
	//input: miniponi
	//output:miniponi!!!!!!!!
	fmt.Println(banger())
	fmt.Println(bangerUni())
}

func banger() string {
	msg := os.Args[1]
	s := msg + strings.Repeat("!", len(msg))

	return s
}
func bangerUni() string {
	msg := os.Args[1]
	l := utf8.RuneCountInString(msg)
	s := msg + strings.Repeat("!", l)

	return s
}
