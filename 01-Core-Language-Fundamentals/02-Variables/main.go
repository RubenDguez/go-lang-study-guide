package main

import "fmt"

func main() {
	fmt.Println("Variables")

	var greeting string // zero-value declaration empty string ""
	greeting = "Hello there"
	fmt.Println(greeting)

	var count int
	count = 10
	fmt.Println(count)

	var isRunning bool
	isRunning = true
	fmt.Println(isRunning)

	var firstName, lastName string
	firstName = "John"
	lastName = "Doe"
	fmt.Printf("%s %s\n", firstName, lastName)

	// type inference

	var year = 2025
	fmt.Println(year)

	// short variable declaration
	email := "test@email.com"
	fmt.Println(email)

	age := 24
	fmt.Println(age)

	isAdult := true
	fmt.Println(isAdult)

}
