package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	in := os.Args
	s := os.Args[1] //strings.Join(in, "")
	n, err := strconv.Atoi(s)

	if len(in) != 2 {
		fmt.Println("Pick a number")
		return
	}

	if err != nil {
		fmt.Printf(`"%v" is not a number`, s)
		return
	}

	if n%2 == 0 {
		fmt.Printf("%v is an even number ", n)
		if n%8 == 0 {
			fmt.Println(" and it is divisible by 8")
		}
	} else {
		fmt.Printf("%v is an odd number", n)
	}
}
