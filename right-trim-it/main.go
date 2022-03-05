package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "tunyacsap"
	fmt.Println(rightTrimIt(name))
}

func rightTrimIt(name string) string {
	s := strings.TrimRight(name, " ")
	return s
}
