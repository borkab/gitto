package main

import (
	"fmt"
	"strings"
)

func main() {
	msg := `
	
	The weather looks good.
I should go and play.
	`

	fmt.Println(trimIt(msg))
}

func trimIt(msg string) string {
	message := strings.TrimSpace(msg)
	return message
}
