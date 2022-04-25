package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	in := os.Args
	s := os.Args[1]
	year, err := strconv.Atoi(s)

	if len(in) != 2 {
		fmt.Println("Give me a year number")
		return
	}
	if err != nil {
		fmt.Printf(`"%v" is not a valid year`, s)
		return
	}
	//if year%4 == 0 && year%100 != 0 || year%400 == 0 {

	var leap bool
	if year%400 == 0 {
		leap = true
	} else if year%100 == 0 {
		leap = false
	} else if year%4 == 0 {
		leap = true
	}

	if leap {
		fmt.Printf("%d is a leap year", year)
	} else {
		fmt.Printf("%d is not a leap year", year)
	}
}
