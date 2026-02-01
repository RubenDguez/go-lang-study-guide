package model

import (
	"errors"
	"time"
)

var EmployeeErrorNilRef = errors.New("nil reference")
var EmployeeErrorInactive = errors.New("employee already inactive")

type Employee struct {
	id        int
	firstName string
	lastName  string
	position  string
	salary    float64
	isActive  bool
	joinedAt  time.Time
}

// Value receiver
func (e Employee) New() (emp *Employee) {
	e.id = 0
	e.firstName = "unknown"
	e.lastName = "unknown"
	e.position = "unknown"
	e.salary = 0.00
	e.isActive = false
	e.joinedAt = time.Now()

	return &e
}

// Value receiver
func (e Employee) Create(id int, firstName string, lastName string, position string, salary float64, isActive bool, joinedAt time.Time) (emp *Employee) {
	e.id = id
	e.firstName = firstName
	e.lastName = lastName
	e.position = position
	e.salary = salary
	e.isActive = isActive
	e.joinedAt = joinedAt

	return &e
}

func (e *Employee) GetFullName() (fullName string, err error) {
	if e == nil {
		return "", EmployeeErrorNilRef
	}

	fullName = e.firstName + " " + e.lastName
	return
}

// Reference receiver
func (e *Employee) Deactivate() (success bool, err error) {
	if e == nil {
		return false, EmployeeErrorNilRef
	}
	if e.isActive == false {
		return false, EmployeeErrorInactive
	}

	e.isActive = false;

	return true, nil
}
