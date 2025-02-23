package main

import (
	"fmt"
	"math"
)

const inflationRate float64 = 2.5

func main() {
	var investmentAmount, years, expectedReturnRate float64

	outputText("Enter investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter Expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter number of years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount,
		expectedReturnRate, years)

	fmt.Printf("Future value: %.2f\n", futureValue)
	fmt.Printf("Future real value: %.2f\n", futureRealValue)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {

	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	frv := fv / math.Pow(1+inflationRate/100, years)
	return fv, frv
}
