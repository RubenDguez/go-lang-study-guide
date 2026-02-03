package main

import "fmt"

type Employee struct {
	id int
	name string
}

func (e Employee) String() string {
	return fmt.Sprintf("Employee[Id: %d | Name: %s]", e.id, e.name)
}

func main() {
	ruben := Employee{
		id: 1,
		name: "Ruben Dominguez",
	}
	fmt.Println(ruben)
}
