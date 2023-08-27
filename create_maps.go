package main

import "fmt"

func main() {
	orders := make(map[string]float32)
	fmt.Println(orders)

	donuts := map[string]int{
		"frosted":   10,
		"chocolate": 15,
		"jelly":     8,
	}
	fmt.Println(donuts)
}
