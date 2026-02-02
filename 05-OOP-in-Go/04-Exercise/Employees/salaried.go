package Employees

import "fmt"

type Salaried struct {
	Name   string
	Salary float64
}

func (s Salaried) CalculatePay() float64 {
	return s.Salary / 12
}

func (s Salaried) String() string {
	return fmt.Sprintf("%-80s $%10.2f", s.Name, s.CalculatePay())
}
