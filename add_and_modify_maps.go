package main

import "fmt"

func main() {
	donuts := map[string]int{
		"frosted":   10,
		"chocolate": 15,
		"jelly":     8,
	}

	// Print out all the donuts
	fmt.Println(donuts)

	donuts["glazed"] = 12
	fmt.Println(donuts["glazed"])

	donuts["jelly"] = 3
	fmt.Println(donuts)
}
