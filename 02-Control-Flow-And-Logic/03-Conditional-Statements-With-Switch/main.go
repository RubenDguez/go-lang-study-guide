package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Conditional Statements with Switch")

	day := "Friday"
	switch day {
	case "Monday", "Tuesday":
		fmt.Println("Start of the week, lots of meetings")
	case "Wednesday", "Thursday":
		fmt.Println("Mid-Week")
	case "Friday":
		fmt.Println("Almost there pal")
	case "Saturday", "Sunday":
		fmt.Println("Weekend, no work!!!")
	default:
		fmt.Println("Unknown weekday")
	}

	switch hour := time.Now().Hour(); {
	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}

	checkType := func(i any) {
		switch v := i.(type) {
		case int:
			fmt.Println("Integer: ", v)
		case bool:
			fmt.Println("Boolean: ", v)
		case string:
			fmt.Println("String: ", v)
		default:
			fmt.Println("Unknown", v)
		}
	}

	checkType(1234)
	checkType(true)
	checkType("Hello World")
	checkType(3.14159)
}
