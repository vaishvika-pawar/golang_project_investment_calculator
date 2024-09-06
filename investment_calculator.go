package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	investmentAmount := outputText("Investment Amount: ")
	expectedFutureRate := outputText("Expected Future Rate: ")
	years := outputText("Years: ")

	futureValue, realFutureValue := calculateFutureValue(investmentAmount, expectedFutureRate, years)

	fmt.Printf("Future value: %.1f\nReal future value: %.1f", futureValue, realFutureValue)
}

func outputText(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}

func calculateFutureValue(investmentAmount, expectedFutureRate, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1+expectedFutureRate/100, years)
	rfv := fv * math.Pow(1+inflationRate/100, years)
	return fv, rfv
}
