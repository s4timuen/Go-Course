package user

import (
	"errors"
	"fmt"
	"time"
)

// casing for private/public
// applies to fields as well
type User struct {
	firstnName string
	lastName   string
	birthdate  string
	createdAt  time.Time // nested
}

// constructor/creator pattern (utility function)
// convention: starts with new
func New(firstName string, lastName string, birthdate string) (*User, error) {

	// validation
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("first name, last name and birthdate are required")
	}

	return &User{ // could be a pointer (for large objects)
		firstnName: firstName,
		lastName:   lastName,
		birthdate:  birthdate,
		createdAt:  time.Now(),
	}, nil
}

// attached to user struct
// with receiver argument -> works like normal argument -> data is copied
func (u *User) PrintUserData() {
	fmt.Println(u.firstnName, u.lastName, u.birthdate, u.createdAt)
}

func (u *User) ClearUserName() {
	u.firstnName = ""
	u.lastName = ""
}
