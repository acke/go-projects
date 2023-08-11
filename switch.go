package main

import "fmt"

func main() {
	name := "H. J. Simp"

	switch name {
    case "Butch": 
      fmt.Println("Head to Robbers Roost.")
    case "Bonnie":
      fmt.Println("Stay put in Joplin.")
    default:
      fmt.Println("Just hide!")
  }
	
}

