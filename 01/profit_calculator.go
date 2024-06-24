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

	fmt.Printf("Earnings: %.2f \n", earnings) // %v %.2f %T %% etc.
	// formattedProfit := fmt.Sprintf("Profit: %.2f \n", profit)
	fmt.Printf("Profit: %.2f \n", profit)
	fmt.Printf("Ratio: %.2f", ratio)
}
