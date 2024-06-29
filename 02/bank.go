package main

import (
	"fmt"
)

func main() {
	var balance float64 = 2000
	var option uint

	// continue keyword to restart loop immediately
	for option != 4 {
		fmt.Print("\n----------\nWelcome to Go Bank!\nWhat do you want to do?\n1. Check balance\n2. Deposit money\n3. Withdraw money\n4. Exit\nYour choice: ")
		fmt.Scan(&option)

		switch option {
		case 1:
			printBalance(balance)
		case 2:
			balance = depositMoney(balance)
		case 3:
			balance = withdrawMoney(balance)
		case 4:
			// do nothing -> exit
		default:
			printInvalidOption()
		}
	}
}

func printBalance(balance float64) {
	fmt.Printf("Your balance is %0.2f", balance)
}

func depositMoney(balance float64) (newBalance float64) {
	var amount float64
	fmt.Print("Desposit amount: ")
	fmt.Scan(&amount)

	if amount < 0 {
		newBalance = balance
		printInvalidAmount("You can not deposit an negativ amount")
	} else {
		newBalance = balance + amount
		printBalance(newBalance)
	}
	return
}

func withdrawMoney(balance float64) (newBalance float64) {
	var amount float64
	fmt.Print("Withdraw amount: ")
	fmt.Scan(&amount)

	if amount > balance || amount < 0 {
		newBalance = balance
		if amount > balance {
			printInvalidAmount("You can not withdraw more than your balance")
		} else if amount < 0 {
			printInvalidAmount("You can not withdraw a negative amount")
		}
	} else {
		newBalance = balance - amount
		printBalance(newBalance)
	}
	return
}

func printInvalidOption() {
	fmt.Print("Invalid option, try again")
}

func printInvalidAmount(text string) {
	fmt.Printf("Invalid Amount: %v", text)
}
