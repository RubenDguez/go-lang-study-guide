package structs

import "errors"

type User struct {
	firstName string
	lastName string
}

func (User) New() (user User) {
	user.firstName = "unknown"
	user.lastName = "unknown"
	return
}

func (u *User) SetFirstName(str string) (err error) {
	if str == "" {
		return errors.New("First name should not be empty")
	}
	u.firstName = str
	return nil
}

func (u *User) SetLastName(str string) (err error) {
	if str == "" {
		return errors.New("Last name should not be empty")
	}
	u.lastName = str
	return nil
}
