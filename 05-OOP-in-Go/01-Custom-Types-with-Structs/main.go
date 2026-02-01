package main

import (
	"fmt"
	"time"

	"github.com/rubendguez/model"
)

func init() {
	fmt.Println("Custom types with Structs")
}

func main() {
	unknown := model.Employee{}.New()
	ruben := model.Employee{}.Create(
		1,
		"Ruben",
		"Dominguez",
		"IT Programmer Analyst",
		100.00,
		true,
		time.Now(),
	)

	fmt.Println(unknown)
	fmt.Println(ruben)
}
