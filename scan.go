package main

import "fmt"

func main() {
	fmt.Println("What would you like for lunch?")

	var food string

	fmt.Scan(&food)

	fmt.Printf("Sure, we can have %v for lunch.", food)
}
