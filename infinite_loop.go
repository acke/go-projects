package main

import (
	"fmt"
)

var number int

func count() {
	var input int
	fmt.Print("Number of guests: ")
	fmt.Scan(&input)
	number = number + input
	fmt.Println("Total guests:", number)
}

func main() {

	// INFINITE LOOP CONTAINING THE COUNT FUNCTION BELOW
	for {
		count()
	}
}
