package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 6.5
	var investmentAmount float64
	expectedReturnRate := 5.5
	var years float64

	fmt.Print("Enter the investment amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Println("Enter the number of years: ")
	fmt.Scan(&years)

	var futureValue = float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
