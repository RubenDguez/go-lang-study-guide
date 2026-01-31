package main

import "fmt"

func main() {
	fmt.Println("Slices")

	items := []int{}
	fmt.Printf("Items: %+v [len: %d cap: %d]\n", items, len(items), cap(items))
	for idx := range 10 {
		items = append(items, idx)
		fmt.Printf("Items: %+v [len: %d cap: %d]\n", items, len(items), cap(items))	
	}

	names := []string{"Ruben", "Lisa", "Alex"}
	fmt.Printf("Names: %+v [len: %d cap: %d]\n", names, len(names), cap(names))

	numbers := make([]int, 3, 10)
	fmt.Printf("Numbers: %+v [len: %d cap: %d]\n", numbers, len(numbers), cap(numbers))

	names = append(names, "Mark")
	fmt.Printf("Names: %+v [len: %d cap: %d]\n", names, len(names), cap(names))

	names = names[:len(names ) - 1]
	fmt.Printf("Names: %+v [len: %d cap: %d]\n", names, len(names), cap(names))
}
