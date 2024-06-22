package main

import (
	"fmt"
)

func CalculateProfit() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	earnings := revenue - expenses
	profit := earnings * (1 - taxRate/100)
	ratio := earnings / profit

	fmt.Print("Earnings: ")
	fmt.Printf("%.2f", earnings)
	fmt.Println("")
	fmt.Print("Profit: ")
	fmt.Printf("%.2f", profit)
	fmt.Println("")
	fmt.Print("Ratio: ")
	fmt.Printf("%.2f", ratio)
}
