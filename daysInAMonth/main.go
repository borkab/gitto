package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	in := os.Args
	month := strings.ToLower(os.Args[1])

	if len(in) != 2 {
		fmt.Println("Give me a month name")
	}

	longmonths := []string{"january", "march", "may", "july", "august", "october", "december"}
	shortmonths := []string{"april", "june", "september", "november"}
	february := []string{"february"}
	allmonths := append(append(append([]string{}, longmonths...), shortmonths...), february...)
	//fmt.Println(allmonths)

	if !contains(month, allmonths) {
		fmt.Printf("%v is not a month", month)
		return
	}
	if contains(month, longmonths) {
		fmt.Printf("%s has 31 days", month)
	} else if contains(month, shortmonths) {
		fmt.Printf("%s has 30 days", month)
	} else if contains(month, february) {
		t := time.Now()
		year := t.Year()
		if year%4 == 0 && year%100 != 0 || year%400 == 0 {
			fmt.Printf("%s has 29 days", month)
		}
		fmt.Printf("%s has 28 days", month)
	}

}

func contains(s string, m []string) bool {

	for _, b := range m {
		if b == s {
			return true
		}
	}
	return false
}
