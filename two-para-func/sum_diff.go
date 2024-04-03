package main

import "fmt"

func sumDiff(num1 int, num2 int) (int, int) {
	sum := num1 + num2
	diff := num1 - num2

	return sum, diff
}

func main() {
	a := 10
	b := 5

	sum, diff := sumDiff(a, b)

	fmt.Printf("Sum: %d, Difference: %d", sum, diff)
}
