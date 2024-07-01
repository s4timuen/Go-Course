package main

import (
	"fmt"

	"example.com/bank/comm"
	"example.com/bank/files"
)

const fileName = "balance.txt"

func main() {
	var balance, err = files.ReadFloatFromFile(fileName)
	var option uint

	if err != nil {
		fmt.Println("\n----------\nERROR")
		fmt.Println(err)
		panic(err)
	}

	for option != 4 {
		comm.PrintOptions()
		fmt.Scan(&option)

		switch option {
		case 1:
			comm.PrintBalance(balance)
		case 2:
			balance = depositMoney(balance)
			files.WriteFloatToFile(balance, fileName)
		case 3:
			balance = withdrawMoney(balance)
			files.WriteFloatToFile(balance, fileName)
		case 4:
			// do nothing -> exit
		default:
			comm.PrintInvalidOption()
		}
	}
}

func depositMoney(balance float64) (newBalance float64) {
	var amount float64
	fmt.Print("Desposit amount: ")
	fmt.Scan(&amount)

	if amount < 0 {
		newBalance = balance
		comm.PrintInvalidAmount("You can not deposit an negativ amount")
	} else {
		newBalance = balance + amount
		comm.PrintBalance(newBalance)
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
			comm.PrintInvalidAmount("You can not withdraw more than your balance")
		} else if amount < 0 {
			comm.PrintInvalidAmount("You can not withdraw a negative amount")
		}
	} else {
		newBalance = balance - amount
		comm.PrintBalance(newBalance)
	}
	return
}
