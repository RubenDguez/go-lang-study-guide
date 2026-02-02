package Employees

import "fmt"

type Hourly struct {
	Name        string
	HourlyRate  float64
	HoursWorked int
}

func (h Hourly) CalculatePay() float64 {
	return h.HourlyRate * float64(h.HoursWorked)
}

func (h Hourly) String() string {
	return fmt.Sprintf("%-80s $%10.2f", h.Name, h.CalculatePay())
}
