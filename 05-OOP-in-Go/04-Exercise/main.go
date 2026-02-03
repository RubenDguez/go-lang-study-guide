package main

import (
	"fmt"

	"github.com/rubendguez/Employees"
)

type IPrintEmployees interface {
	Employees.Salaried | Employees.Commission | Employees.Hourly
	CalculatePay() float64
	String() string
}

var salariedEmployees = make([]Employees.Salaried, 0)
var commissionEmployees = make([]Employees.Commission, 0)
var hourlyEmployees = make([]Employees.Hourly, 0)

func PrintEmployees[T IPrintEmployees](emps []T) (total float64) {
	for _, emp := range emps {
		fmt.Println(emp)
		total += emp.CalculatePay()
	}
	return
}

func init() {
	fmt.Println("Exercise: Payroll Processor")
}

func main() {
	fmt.Println("Welcome to the Payroll Processor!")

	ruben := Employees.Salaried{
		Name:   "Ruben Dominguez",
		Salary: 100000.00,
	}
	salariedEmployees = append(salariedEmployees, ruben)

	lisa := Employees.Commission{
		Name:           "Lisa Freiwald",
		BaseSalary:     30000.00,
		Sales:          80000.00,
		CommissionRate: 0.05,
	}
	commissionEmployees = append(commissionEmployees, lisa)

	alex := Employees.Hourly{
		Name:        "Alexander Ramhit",
		HourlyRate:  12.00,
		HoursWorked: 120,
	}
	hourlyEmployees = append(hourlyEmployees, alex)


	payroll := 0.00

	fmt.Println()
	payroll += PrintEmployees(salariedEmployees)
	payroll += PrintEmployees(commissionEmployees)
	payroll += PrintEmployees(hourlyEmployees)

	fmt.Printf("\nTotal Payroll: $%.2f\n", payroll)
}
