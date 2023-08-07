package main
import "fmt"

func main() {
    myTutors := []string{"Kirsty", "Mishell", "Jose", "Neil"}
    changeLastElement(myTutors, "Bobby")
}

func changeLastElement(myTutors []string, element string) {
  if (len(myTutors) > 0) {
    myTutors[len(myTutors)-1] = element
   }
   fmt.Println(myTutors)
}

