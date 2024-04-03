package main

import "fmt"

func main() {
	var numberOfDays int

	fmt.Print("Enter the number of days: ")
	fmt.Scan(&numberOfDays)

	switch {
	case 28 == numberOfDays:
		fmt.Println("It's February")
	default:
		fmt.Println("It's not February")
	}
}
