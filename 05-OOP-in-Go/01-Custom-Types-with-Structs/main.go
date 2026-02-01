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

	fmt.Printf("%+v\n", *unknown)
	fmt.Printf("%+v\n", *ruben)
	ruben.Deactivate()
	fmt.Printf("%+v\n", *ruben)

	_, err := ruben.Deactivate()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("User deactivated")
		fmt.Printf("%+v\n", *ruben)
	}

	fmt.Println(ruben.GetFullName())
}
