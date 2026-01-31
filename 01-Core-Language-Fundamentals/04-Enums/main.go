package main

import "fmt"

type LogLevel int
const (
	Debug LogLevel = iota
	Info
	Warn
	Error
	Fatal
)

const (
	Monday  = "Monday"
	Tuesday = "Tuesday"
	Wednesday = "Wednesday"
	Thursday = "Thursday"
	Friday = "Friday"
	Saturday = "Saturday"
	Sunday = "Sunday"
)

func main() {
	fmt.Println("Enums")

	fmt.Println(Monday)
	fmt.Println(Tuesday)
	fmt.Println(Wednesday)
	fmt.Println(Thursday)
	fmt.Println(Friday)
	fmt.Println(Saturday)
	fmt.Println(Sunday)

	fmt.Println(Debug)
	fmt.Println(Info)
	fmt.Println(Warn)
	fmt.Println(Error)
	fmt.Println(Fatal)
}
