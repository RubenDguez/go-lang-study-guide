package main

import (
	"fmt"
	"os"
)

func Defer() {
	fmt.Println("Statement 1")
	fmt.Println("Statement 2")
	defer fmt.Println("Deferred statement 1")
	fmt.Println("Statement 3")
}

func main() {
	fmt.Println("Defer")
	Defer()

	defer func() {
		fmt.Println("Statement 1")
	}()

	file, err := os.Create("./text.txt")
	if err != nil {
		fmt.Println("Error", err)
	}
	defer func() {
		fmt.Println("Closing file stream")
		file.Close()
	}()

	fmt.Println("Statement 2")
	fmt.Println("Statement 3")
}
