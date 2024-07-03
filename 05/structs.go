package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	firstnName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (DD/MM/YYYY): ")

	appUser, err := user.New(
		firstnName,
		lastName,
		birthdate,
	)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		return
	}

	// printUserData(&appUser)
	appUser.PrintUserData()
	appUser.ClearUserName()
	appUser.PrintUserData()

	email := getUserData("Please enter your email: ")
	password := getUserData("Please enter your password: ")

	adminUser, err := user.NewAdmin(
		firstnName,
		lastName,
		birthdate,
		email,
		password,
	)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		return
	}

	adminUser.User.PrintUserData()
	adminUser.PrintAdminData()
}

func getUserData(promptText string) (value string) {
	fmt.Print(promptText)
	fmt.Scanln(&value)
	return
}

/*
func printUserData(user *user) {
	fmt.Println((*user).firstnName, user.lastName, user.birthdate, user.createdAt) // deref and shortcut
}
*/
