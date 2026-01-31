package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println("More Slices")
	fmt.Println("Advanced sliced operations")

	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("Numbers: %+v [len: %d] [cap: %d]\n", numbers, len(numbers), cap(numbers))

	s1 := numbers[2:5] // low:high (high is exclusive - not included)
	fmt.Printf("s1: %+v [len: %d] [cap: %d]\n", s1, len(s1), cap(s1))

	s2 := numbers[:5] // low:high (high is exclusive - not included)
	fmt.Printf("s2: %+v [len: %d] [cap: %d]\n", s2, len(s2), cap(s2))

	s3 := numbers[3:] // low:high (high is exclusive - not included) high is the len() in this case
	fmt.Printf("s3: %+v [len: %d] [cap: %d]\n", s3, len(s3), cap(s3))

	s4 := numbers[:] // clone of the original one
	fmt.Printf("s4: %+v [len: %d] [cap: %d]\n", s4, len(s4), cap(s4))

	numberToSearch := 8
	if contains := slices.Contains(s4, numberToSearch); contains {
		fmt.Println("The slice contains ", numberToSearch)
	}


	greetings := "Hello there!!!"
	ss1 := greetings[5:]
	fmt.Printf("ss1: %s [len: %d]\n", ss1, len(ss1))
}
