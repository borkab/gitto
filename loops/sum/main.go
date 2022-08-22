package main

import "fmt"

func main() {
	y := 0
	fmt.Println(sum(y))
}

func sum(x int) int {

	for i := 0; i <= 10; i++ {
		x += i
	}
	return x
}
