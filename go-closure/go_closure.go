package main

import "fmt"

func greet(name string) {
	greetHello := func() {
		fmt.Printf("Hello %v", name)
	}
	greetHello()
}

func main() {
	greet("Azam")
}
