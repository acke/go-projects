package main

import "fmt"

type Point struct {
	x int
	y int
}

type Circle struct {
	Point
	radius int
}

func main() {
	circle := Circle{Point{4, 5}, 2}

	fmt.Println(circle.x)

	fmt.Println(circle)
}
