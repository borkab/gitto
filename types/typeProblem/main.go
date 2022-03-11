package main

import "fmt"

/*exercise: solve the data type problem in the program

expected output:
width: 265 height: 265
are they equal? true
*/

func main() {
	//change the  following data types to the correct
	//data types where appropriate.
	var (
		width  uint16
		height uint16
	)

	//don't touch this:
	width, height = 255, 265
	width += 10
	fmt.Printf("width: %d height: %d\n", width, height)
	fmt.Println("are they equal?", width == height)
}
