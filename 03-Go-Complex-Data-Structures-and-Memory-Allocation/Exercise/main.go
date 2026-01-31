package main

import (
	"errors"
	"fmt"
	"slices"

	"github.com/rubendguez/structs"
)

var contacts = []structs.Contact{}

func Push(s *[]structs.Contact, c structs.Contact) (err error) {
	if c.GetFirstName() == c.GetLastName() && c.GetFirstName() == "" {
		return errors.New("First name nor Last name fields should not be empty")
	}

	*s = append(*s, c)
	return
}

func Pop(s *[]structs.Contact) (c *structs.Contact, err error) {
	if len(*s) == 0 {
		return c, errors.New("The slice is already empty, nothing to pop")
	}

	c = &(*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return
}

func Shift(s *[]structs.Contact, c structs.Contact) (err error) {
	if c.GetFirstName() == c.GetLastName() && c.GetFirstName() == "" {
		return errors.New("First name nor Last name fields should not be empty")
	}

	*s = slices.Insert((*s), 0, c)

	return
}

func Unshift(s *[]structs.Contact) (c *structs.Contact, err error) {
	if len(*s) == 0 {
		return c, errors.New("The slice is already empty, nothing to pop")
	}

	c = &(*s)[0]
	*s = (*s)[1:]

	return
}

func GetByFirstName(s *[]structs.Contact, str string) (c *structs.Contact, err error) {
	for _, contact := range *s {
		if contact.GetFirstName() == str {
			return &contact, nil
		}
	}
	return c, errors.New("No contact found with first name " + str)
}

func GetByLastName(s *[]structs.Contact, str string) (c *structs.Contact, err error) {
	for _, contact := range *s {
		if contact.GetLastName() == str {
			return &contact, nil
		}
	}
	return c, errors.New("No contact found with last name " + str)
}

func Print(s *[]structs.Contact) {
	for idx, contact := range *s {
		fmt.Println(idx, contact.ToString())
	}
}

func main() {
	fmt.Println("Contact Management System")
	c1 := structs.Contact{}.New()
	c2 := structs.Contact{}.New()
	c3 := structs.Contact{}.New()
	c4 := structs.Contact{}.New()

	c1.SetFirstName("Ruben")
	c1.SetLastName("Dominguez")
	c1.SetPhoneNumber("1234567890")

	c2.SetFirstName("Lisa")
	c2.SetLastName("Freiwald")
	c2.SetPhoneNumber("0987654321")

	c3.SetFirstName("Alexander")
	c3.SetLastName("Ramhit")
	c3.SetPhoneNumber("6789012345")

	c4.SetFirstName("Samantha")
	c4.SetLastName("Dominguez")
	c4.SetPhoneNumber("8099012345")

	fmt.Println("\nAdding some contacts with Push")
	Push(&contacts, c1)
	fmt.Println("Added contact: ", c1.ToString())
	Push(&contacts, c2)
	fmt.Println("Added contact: ", c2.ToString())
	Push(&contacts, c3)
	fmt.Println("Added contact: ", c3.ToString())

	fmt.Println("\nRemoving some contacts with Pop")
	if contact, err := Pop(&contacts); err == nil {
		fmt.Println("Removed contact", contact.ToString())
	}

	if contact, err := Unshift(&contacts); err == nil {
		fmt.Println("Removed contact", contact.ToString())
	}

	fmt.Println("\nAdding some contacts with Shift")
	Shift(&contacts, c4)
	fmt.Println("Added contact: ", c4.ToString())


	fmt.Println("\nPrint current contacts")
	Print(&contacts)

	fmt.Println("\nGet By")
	if contact, err := GetByFirstName(&contacts, "Lisa"); err == nil {
		fmt.Println("Get by first name \"Lisa\": ", contact.ToString())
	}
	if contact, err := GetByLastName(&contacts, "Dominguez"); err == nil {
		fmt.Println("Get by last name \"Lisa\": ", contact.ToString())
	}
}
