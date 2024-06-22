package main

import (
	"fmt"
	"math"
)

func CalculateInvestment() {
	const inflationRate = 2.5
	var investmentAmount float64 = 1000
	expectedReturnRate := 5.5
	var years float64

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)
	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Print("Future Value: ")
	fmt.Printf("%.2f", futureValue)
	fmt.Println("")
	fmt.Print("Future Real Value: ")
	fmt.Printf("%.2f", futureRealValue)
}