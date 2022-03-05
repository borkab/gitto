package main

import (
	"fmt"
	"strings"
)

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
