package model

import "time"

type Employee struct {
	id        int
	firstName string
	lastName  string
	position  string
	salary    float64
	isActive  bool
	joinedAt  time.Time
}

func (Employee) New() (emp Employee) {
	emp.id = 0
	emp.firstName = "unknown"
	emp.lastName = "unknown"
	emp.position = "unknown"
	emp.salary = 0.00
	emp.isActive = false
	emp.joinedAt = time.Now()

	return
}

func (Employee) Create(id int, firstName string, lastName string, position string, salary float64, isActive bool, joinedAt time.Time) (emp Employee) {
	emp.id = id
	emp.firstName = firstName
	emp.lastName = lastName
	emp.position = position
	emp.salary = salary
	emp.isActive = isActive
	emp.joinedAt = joinedAt

	return
}
