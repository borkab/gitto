package main

import "fmt"

func main() {
	var x int
	for i := 0; i <= 10; i++ {
		x += i
	}
	fmt.Println(x)
}
