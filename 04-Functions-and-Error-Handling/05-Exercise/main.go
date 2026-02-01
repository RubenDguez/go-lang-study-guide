package main

import (
	"fmt"
	"strings"
)

const (
	division             = "Division"
	divisionErrorMessage = "division by zero"
)

type MathError struct {
	Operation string
	InputA    float64
	InputB    float64
	Message   string
}

func (e MathError) Error() string {
	var inputs []string

	if e.Operation == division {
		inputs = append(inputs, fmt.Sprintf("a=%f", e.InputA))
		inputs = append(inputs, fmt.Sprintf("b=%f", e.InputB))
	}

	return fmt.Sprintf("Math error in %s (%s): %s", e.Operation, strings.Join(inputs, ", "), e.Message)
}

func Sum(numbers ...int) (total int, err error) {
	defer fmt.Println("Sum finished...")

	for _, val := range numbers {
		total += val
	}

	return
}

func Division(a, b float64) (total float64, err error) {
	if b == 0 {
		return 0, &MathError{
			Operation: division,
			InputA:    a,
			InputB:    b,
			Message:   divisionErrorMessage,
		}
	}

	total = a / b
	return
}

func init() {
	fmt.Println("Functions and Errors Handling Exercise")
}

func main() {
	if total, err := Division(4.0, 1.3); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Total:", total)
	}

	if total, err := Division(4.0, 0.00); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Total:", total)
	}
}
