package main

import (
	"fmt"
)

func CalculateProfit() {
	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")
	taxRate := getUserInput("Tax Rate: ")
	earnings, profit, ratio := calculateFinancials(revenue, expenses, taxRate)
	printFinancials(earnings, profit, ratio)
}

func getUserInput(text string) (inputValue float64) {
	fmt.Print(text)
	fmt.Scan(&inputValue)
	return
}

func calculateFinancials(revenue, expenses, taxRate float64) (earnings, profit, ratio float64) {
	earnings = revenue - expenses
	profit = earnings * (1 - taxRate/100)
	ratio = earnings / profit
	return
}

func printFinancials(earnings, profit, ratio float64) {
	fmt.Printf("Earnings: %.2f \n", earnings) // %v %.2f %T %% etc.
	// formattedProfit := fmt.Sprintf("Profit: %.2f \n", profit)
	fmt.Printf("Profit: %.2f \n", profit)
	fmt.Printf("Ratio: %.2f", ratio)
}
