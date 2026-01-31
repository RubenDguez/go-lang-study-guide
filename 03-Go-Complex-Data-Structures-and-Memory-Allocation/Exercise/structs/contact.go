package structs

import (
	"errors"
	"fmt"
)

type Contact struct {
	firstName   string
	lastName    string
	phoneNumber string
}

func (Contact) New() Contact {
	return Contact{
		firstName:   "unknown",
		lastName:    "unknown",
		phoneNumber: "unknown",
	}
}

func (current *Contact) GetFirstName() (string) {
	return current.firstName
}

func (current *Contact) SetFirstName(str string) (err error) {
	if str == "" {
		return errors.New("First name should not be empty")
	}

	current.firstName = str
	return
}

func (current *Contact) GetLastName() (string) {
	return current.lastName
}

func (current *Contact) SetLastName(str string) (err error) {
	if str == "" {
		return errors.New("Last name should not be empty")
	}

	current.lastName = str
	return 
}

func (current *Contact) SetPhoneNumber(str string) (err error) {
	if str == "" {
		return errors.New("Phone number should not be empty")
	}

	current.phoneNumber = str
	return
}

func (current *Contact) Update(c Contact) (err error) {
	if c.firstName == "" {
		return errors.New("First name should not be empty")
	}
	if c.lastName == "" {
		return errors.New("Last name should not be empty")
	}

	current.firstName = c.firstName
	current.lastName = c.lastName
	current.phoneNumber = c.phoneNumber

	return
}

func (current *Contact) ToString() string {
	return fmt.Sprintf("[First name: %s] [Last name: %s] [Phone number: %s]", current.firstName, current.lastName, current.phoneNumber)
}
