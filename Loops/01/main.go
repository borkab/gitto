package main

import "fmt"

// EXERCISE: Sum the Numbers
//
//  1. By using a loop, sum the numbers between 1 and 10.
//  2. Print the sum.
//
// EXPECTED OUTPUT
//
//	Sum: 55
func main() {
	a := 0
	fmt.Println(sum(a))

}

func sum(a int) int {
	for i := 0; i <= 10; i++ {
		a += i
	}
	return a
}
