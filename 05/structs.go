package main

import (
	"fmt"
	"time"
)

// casing for private/public
type user struct {
	firstnName string
	lastName   string
	birthdate  string
	createdAt  time.Time // nested
}

// attached to user struct
// with receiver argument -> works like normal argument -> data is copied
func (u *user) printUserData() {
	fmt.Println(u.firstnName, u.lastName, u.birthdate, u.createdAt)
}

func (u *user) clearUserName() {
	u.firstnName = ""
	u.lastName = ""
}

func main() {
	var appUser = user{
		firstnName: getUserData("Please enter your first name: "),
		lastName:   getUserData("Please enter your last name: "),
		birthdate:  getUserData("Please enter your birthdate (DD/MM/YYYY): "),
		createdAt:  time.Now(),
	}
	// printUserData(&appUser)
	appUser.printUserData()
	appUser.clearUserName()
	appUser.printUserData()
}

func getUserData(promptText string) (value string) {
	fmt.Print(promptText)
	fmt.Scan(&value)
	return
}

/*
func printUserData(user *user) {
	fmt.Println((*user).firstnName, user.lastName, user.birthdate, user.createdAt) // deref and shortcut
}
*/
