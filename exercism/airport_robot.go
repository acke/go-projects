package airportrobot

import "fmt"

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
type Greeter interface {
    LanguageName() string
    Greet(name string) string
}

type Italian struct {
}

func (i Italian) LanguageName() string{
    return "Italian"
}

func (i Italian) Greet(name string) string {
    return fmt.Sprintf("Ciao %v!", name)
}

type Portuguese struct {
}

func (p Portuguese) LanguageName() string{
    return "Portuguese"
}

func (p Portuguese) Greet(name string) string {
    return fmt.Sprintf("Olá %v!", name)
}


func SayHello(name string, greeter Greeter) string{
    return fmt.Sprintf("I can speak %s: %s", greeter.LanguageName(), greeter.Greet(name))
}

