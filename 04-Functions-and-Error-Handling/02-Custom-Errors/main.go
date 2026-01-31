package main

import (
	"errors"
	"fmt"
)

var ErrorDivideByZero = errors.New("division by zero")
var ErrorNumeratorTooLarge = errors.New("numerator is too large")
var ErrorDivideByNegativeNumber = errors.New("division by negative number")

func divide(numerator, denominator float64) (val float64, err error) {
	if denominator == 0 {
		return 0, ErrorDivideByZero 
	}
	if denominator < 0 {
		return 0, ErrorDivideByNegativeNumber
	}
	if numerator > 2000 {
		return 0, ErrorNumeratorTooLarge
	}

	val = numerator / denominator
	return 
}

func main() {
	numerator := 2000.0
	denominator := -2.0

	val, err := divide(numerator, denominator)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println("Result:", val)
}
