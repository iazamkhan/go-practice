package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Enter the yearly revenue")
	fmt.Scan(&revenue)
	fmt.Print("Enter the yearly expenses")
	fmt.Scan(&expenses)
	fmt.Print("Enter the default taxRate")
	fmt.Scan(&taxRate)
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	fmt.Printf("Profit is: %.1f\nProfit ratio is: %.1f", profit, ratio)
	// fmt.Print(formattedValue)
}
