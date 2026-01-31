package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/rubendguez/structs"
)

func add1(v int) (err error) {
	v += 1
	return nil
}

func add2(v *int) (err error) {
		if v == nil {
		return errors.New("Value v should not be nil")
	}

	*v += 2
	return nil
}

func main() {
	fmt.Println("Pointers")
	var user structs.User = structs.User{}.New()
	fmt.Println(user)

	if err := user.SetFirstName("Ruben"); err != nil {
		log.Fatalf("%+v\n", err)
	}

	if err := user.SetLastName("Dominguez"); err != nil {
		log.Fatalf("%+v\n", err)
	}

	fmt.Println(user)

	count := 1
	fmt.Println(count)
	add1(count)
	fmt.Println(count)
	if err := add2(&count); err == nil {
		fmt.Println(count)
	}
}
