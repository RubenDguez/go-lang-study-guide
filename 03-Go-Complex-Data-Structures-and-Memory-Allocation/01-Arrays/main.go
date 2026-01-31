package main

import "fmt"

func main() {
	fmt.Println("Arrays")

	var numbers = [2]int{}
	fmt.Printf("%+v\n", numbers)

	numbers = [2]int{1, 2}
	fmt.Printf("%+v\n", numbers)

	numbers[0] = 3
	numbers[1] = 4
	fmt.Printf("%+v\n", numbers)

	prime := [4]int{2, 3, 5, 7}
	fmt.Printf("%+v\n", prime)

	for idx, val := range prime {
		fmt.Printf("[ %d : %d ] ", idx, val)
	}
	fmt.Println()

	fmt.Printf("Prime array length: %d\n", len(prime))

	var matrix = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Printf("%+v\n", matrix)

	for j, row := range matrix {
		for k, cell := range row {
			fmt.Printf("(%d,%d): %d\n", j, k, cell)
		}
	}
}
