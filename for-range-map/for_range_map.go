package main

import "fmt"

func main() {
	marks := map[string]int{"Java": 85, "Python": 90, "Go": 95}

	fmt.Println("Marks: ")
	for subject, Marks := range marks {
		fmt.Printf("%s: %d\n", subject, Marks)
	}
}
