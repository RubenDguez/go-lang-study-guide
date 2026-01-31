package main

import "fmt"

func main() {
	fmt.Println("Understanding Values")

	var t any

	fmt.Printf("Hello " + "World" + "\n")

	fmt.Printf("%d\n", (1 + 1))

	fmt.Printf("%f\n", 3.14159)

	fmt.Println(true, false)

	fmt.Printf("%+v\n", []int{1,2,3,4,5})

	fmt.Printf("%+v\n", t)
}
