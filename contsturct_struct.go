package main

import "fmt"

type Country struct {
	name      string
	capital   string
	latitude  float32
	longitude float32
}

func main() {
	var france Country
	fmt.Println(france)
}
