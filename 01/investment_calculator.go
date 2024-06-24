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

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, inflationRate, years)

	fmt.Printf("Future Value: %.2f \n", futureValue)
	fmt.Printf("Future Real Value: %.2f", futureRealValue)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, inflationRate, years float64) (futureValue float64, futureRealValue float64) {
	futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue = futureValue / math.Pow(1+inflationRate/100, years)
	// return
	return futureValue, futureRealValue
}
