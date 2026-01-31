package main

import (
	"errors"
	"fmt"
)

func index(start int) func() int {
	return func() int {
		start++
		return start
	}
}

func counter(start int, end int) {
	if start >= end {
		return
	}

	fmt.Println(start)
	next := start + 1
	counter(next, end)
}

func init() {
	fmt.Println("Function")
}

func divide(a, b float64) (val float64, err error) {
	if b <= 0 {
		return val, errors.New("Cannot divide by zero or negative value")
	}

	val = a / b
	return
}

func main() {
	fmt.Println("function as value")
	myfunc := func(a, b int) (total int) {
		return a + b
	}
	fmt.Println(myfunc(1, 2))

	fmt.Println("Closure")
	indexer := index(5)
	fmt.Println(indexer())
	fmt.Println(indexer())
	fmt.Println(indexer())

	fmt.Println("Recursion")
	counter(2, 6)

	fmt.Println("Variadic")
	variadic := func(n ...int) (total int) {
		for _, val := range n {
			total += val
		}

		return
	}
	fmt.Println(variadic(1, 2, 3, 4))

	fmt.Println("Using variadic for optional params")
	config := func(n ...string) {
		if len(n) == 0 {
			fmt.Println("No configuration passed.")
			return
		}

		for _, val := range n {
			fmt.Println(val)
		}
	}

	config()
	config("foo", "bar", "fizz", "buzz")

	fmt.Println("Returning multiple values")
	numerator := 4.0
	denominator := 5.0
	if val, err := divide(numerator, denominator); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Dividing", numerator, "by", denominator, "equals", val)
	}
}
