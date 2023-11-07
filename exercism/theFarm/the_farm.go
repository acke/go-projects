package thefarm

import (
	"errors"
	"fmt"
)

type InvalidCowsError struct {
	numberOfCows int
	message      string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%v cows are invalid: %s", e.numberOfCows, e.message)
}

func DivideFood(fodderCalculator FodderCalculator, numberOfCows int) (float64, error) {
	amountOfFodder, error := fodderCalculator.FodderAmount(numberOfCows)
	if error != nil {
		return 0.0, error
	}

	foodPerCow, error := fodderCalculator.FatteningFactor()
	if error != nil {
		return 0.0, error
	}

	return (amountOfFodder * foodPerCow) / float64(numberOfCows), error
}

func ValidateInputAndDivideFood(fodderCalculator FodderCalculator, numberOfCows int) (float64, error) {
	if numberOfCows > 0 {
		return DivideFood(fodderCalculator, numberOfCows)
	}
	return 0, errors.New("invalid number of cows")
}

func ValidateNumberOfCows(numberOfCows int) error {
	if numberOfCows < 0 {
		return &InvalidCowsError{numberOfCows, "there are no negative cows"}
	}

	if numberOfCows == 0 {
		return &InvalidCowsError{numberOfCows, "no cows don't need food"}
	}

	return nil
}
