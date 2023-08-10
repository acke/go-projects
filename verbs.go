package main

import "fmt"

func main() {
  floatExample := 1.75
  
  fmt.Printf("Working with a %T", floatExample) 
  
  fmt.Println("\n***") // Added for spacing
  
  yearsOfExp := 3
  reqYearsExp := 15

  fmt.Printf("I have %d years of Go experience and this job is asking for %d years.", yearsOfExp, reqYearsExp) 
  
  fmt.Println("\n***") // Added for spacing
  
  stockPrice := 3.50
  
  fmt.Printf("Each share of Gopher feed is $%.2f!", stockPrice) 
}

