package main

import "fmt"

func main() {
	myTutors := []string{"Kirsty", "Mishell", "Jose", "Neil"}
	myTutors = append(myTutors, "Josh")
	fmt.Println(myTutors)
}
