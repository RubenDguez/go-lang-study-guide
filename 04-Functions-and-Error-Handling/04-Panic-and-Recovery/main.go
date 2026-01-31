package main

import "fmt"

func mayPanic(willPanic bool) {
	if willPanic {
		panic("Something went really wrong")
	}

	fmt.Println("Ran without panic")
}

func recoverable() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %+v\n", r)
		}
	}()

	mayPanic(true)
}

func main() {
	mayPanic(false)
	recoverable()
}
