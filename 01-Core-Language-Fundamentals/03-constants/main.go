package main

import "fmt"

const HOST = "localhost"

func main() {
	fmt.Println("Constants")

	fmt.Println(HOST)

	const APP_NAME = "Go"
	fmt.Println(APP_NAME)

	const PI float32 = 3.14159
	fmt.Println(PI)


	const RATE float32 = 5.2
	fmt.Println(RATE)
}
