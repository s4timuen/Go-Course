package comm

import (
	"fmt"

	"github.com/pallinder/go-randomdata"
)

func PrintOptions() {
	fmt.Println("\n----------")
	fmt.Println("Welcome to Go Bank!")
	fmt.Print("Service: ")
	fmt.Println(randomdata.PhoneNumber())
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")
	fmt.Print("Your choice: ")
}

func PrintBalance(balance float64) {
	fmt.Printf("Your balance is %0.2f", balance)
}

func PrintInvalidOption() {
	fmt.Print("Invalid option, try again")
}

func PrintInvalidAmount(text string) {
	fmt.Printf("Invalid Amount: %v", text)
}
