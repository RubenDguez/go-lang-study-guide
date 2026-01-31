package main

import "fmt"

func main() {
	fmt.Println("Maps")

	grades := map[string]int{
		"Ruben": 80,
		"Lisa":  95,
		"Alex":  91,
	}

	fmt.Printf("%+v\n", grades)
	
	ruben, ok := grades["Ruben"]
	fmt.Println(ruben, ok)

	student := "Ruben"
	grade := 90
	if _, ok := grades[student]; ok {
		grades[student] = grade
	}
	fmt.Printf("%+v\n", grades)

	student = "Lisa"
	if val, ok := grades[student]; ok {
		fmt.Printf("%s : %d\n", student, val)
	}

	student = "Alex"
	delete(grades, student)
	fmt.Printf("%+v\n", grades)

	var configs = make(map[string]string)
	fmt.Printf("%+v %T\n", configs, configs)
	

	configs["env"] = "neerqa"
	fmt.Printf("%+v %T\n", configs, configs)
}
