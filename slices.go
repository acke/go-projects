package main

import "fmt"

func main() {
    myTutors := [4]string{"Kirsty", "Mishell", "Jose", "Neil"}
    tutorsSlice := myTutors[:]
    subjects := [] string {"Go", "Java", "Python"}
    fmt.Println(tutorsSlice)
    fmt.Println(subjects)
}
