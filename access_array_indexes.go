package main

import "fmt"

func main() {
	triangleAngles := [3]int{30, 60, 90}
	fmt.Println(triangleAngles[2])
	sum := triangleAngles[0] + triangleAngles[1] + triangleAngles[2]
	fmt.Println(sum)
}
