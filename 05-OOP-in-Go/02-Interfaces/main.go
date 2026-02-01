package main

import "fmt"

type Human interface {
	Eat() (favorite string)
}

type Person interface {
	Human
	GetId() (id int)
	GetName() (name string)
}

type Employee struct {
	Id   int
	Name string
}

type Writer struct {
	Id   int
	Name string
}

func (e *Employee) Eat() (favorite string) {
	return "Steak"
}

func (e *Employee) GetId() (id int) {
	return e.Id
}

func (e *Employee) GetName() (name string) {
	return e.Name
}

func (w *Writer) Eat() (favorite string) {
	return "Vegetables"
}

func (w *Writer) GetId() (id int) {
	return w.Id
}

func (w *Writer) GetName() (name string) {
	return w.Name
}

func displayPerson(p Person) {
	fmt.Println(p.GetId(), p.GetName(), p.Eat())
}

func init() {
	fmt.Println("Interfaces")
}

func main() {
	ruben := Employee{
		Id: 1,
		Name: "Ruben Dominguez",
	}

	lisa := Writer{
		Id: 1,
		Name: "Lisa Freiwald",
	}

	displayPerson(&ruben)
	displayPerson(&lisa)
}
