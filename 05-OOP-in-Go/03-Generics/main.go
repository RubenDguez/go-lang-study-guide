package main

import "fmt"

type Numbers interface {
	int | float32 | float64
}

func Sum[T Numbers](vals ...T) (result T) {
	for _, val := range vals {
		result += val
	}

	return
}

func init() {
	fmt.Println("Generics")
}

func main() {
	fmt.Println("Integers")
	intResult := Sum(1, 2, 3, 4)
	fmt.Println(intResult)

	fmt.Println("Float32")
	float32Result := Sum[float32](1.0, 1.2, 1.3)
	fmt.Println(float32Result)

	fmt.Println("Float64")
	float64Result := Sum(1.0, 1.2, 1.3)
	fmt.Println(float64Result)
}
