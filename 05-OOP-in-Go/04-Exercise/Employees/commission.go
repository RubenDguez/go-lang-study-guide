package Employees

import "fmt"

type Commission struct {
	Name           string
	BaseSalary     float64
	CommissionRate float64
	Sales          float64
}

func (c Commission) CalculatePay() float64 {
	return (c.BaseSalary / 12) + (c.CommissionRate * c.Sales)
}

func (c Commission) String() string {
	return fmt.Sprintf("%-80s $%10.2f", c.Name, c.CalculatePay())
}
