package main

import "fmt"

func main() {
	fmt.Println("Conditional Statements with If - Else")

	fmt.Println("Check temperature")
	temp := 70
	if temp < 70 {
		fmt.Println("Temperature is less than 70 degrees")
	} else {
		fmt.Println("Temperature is not less than 70 degrees")
	}

	fmt.Println("Check Grades")
	grade := 85
	if grade >= 90 {
		fmt.Println("Grade: A")
	} else if grade >= 80 {
		fmt.Println("Grade: B")
	} else if grade >= 70 {
		fmt.Println("Grade: C")
	} else if grade >= 60 {
		fmt.Println("Grade: D")
	} else {
		fmt.Println("Failed")
	}

	access := map[string]bool{
		"Ruben": true,
		"Lisa":  false,
	}

	fmt.Println("If statement with assignment")
	if haveAccess, exists := access["Ruben"]; exists && haveAccess {
		fmt.Println("Access granted")
	} else {
		fmt.Println("Access denied")
	}

	// You could write the previous statement as
	haveAccess, exists := access["Lisa"]
	if exists && haveAccess {
		fmt.Println("Access granted")
	} else {
		fmt.Println("Access denied")
	}

}
