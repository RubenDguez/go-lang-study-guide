package main

import "fmt"

func main() {
	fmt.Println("For Loop")

	fmt.Println("C-Style loop")
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	fmt.Println("While-Style loop")

	idx := 0
	for idx < 3 {
		fmt.Println(idx)
		idx++
	}

	fmt.Println("Range-Style loop")
	for i := range 3 {
		fmt.Println(i)
	}

	fmt.Println("Skipping")
	for i := range 10 {
		if i % 2 == 0 {
			continue
		}

		fmt.Println(i)
	}

	fmt.Println("Iterable")
	items := []string{"Go", "TypeScript", "Java"}
	for idx, val := range items {
		fmt.Printf("%d [ %s ]\n", idx, val)
	}
	fmt.Println("Iterable: Just Index")
	for idx := range items {
		fmt.Println(idx)
	}
	fmt.Println("Iterable: Just Values")
	for _, val := range items {
		fmt.Println(val)
	}
}
