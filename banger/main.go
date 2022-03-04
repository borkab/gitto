package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	msg := os.Args[1]
	fmt.Println(banger(msg))
	fmt.Println(bangerUni(msg))
}

func banger(msg string) string {

	s := msg + strings.Repeat("!", len(msg))

	return s
}
func bangerUni(msg string) string {

	l := utf8.RuneCountInString(msg)
	s := msg + strings.Repeat("!", l)

	return s
}
