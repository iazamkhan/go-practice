package main

import "fmt"

func main() {
	number := 5
	i := 1

	for i <= 10 {
		fmt.Printf("%d * %d = %d\n", number, i, number*i)
		i++
	}
}
