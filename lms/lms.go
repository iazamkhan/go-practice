package main

import "fmt"

func main() {
	type Library struct {
		bookName      string
		authorName    string
		numberOfPages int
	}

	var count, input, i int = 0, 0, 0

	var myLib [20]Library
	// fmt.Print(Book1.authorName)
	fmt.Printf("Welcome to the Library\nEnter 1 to add a book\nEnter 2 to display books information\nEnter 3 to list the count of books in the library\n")
	fmt.Scan(&input)

	switch input {
	case 1:
		fmt.Println("Enter book name = ")
		fmt.Scan(&myLib[i].bookName)
		fmt.Println("Enter author name = ")
		fmt.Scan(&myLib[i].authorName)
		fmt.Println("Enter pages = ")
		fmt.Scan(&myLib[i].numberOfPages)
		count++

	case 2:
		for i = 0; i < count; i++ {
			fmt.Printf("Book name: %v\nAuthor name: %v\nNumber of pages: %v", myLib[i].authorName, myLib[i].bookName, myLib[i].numberOfPages)
		}
	case 3:
		fmt.Printf("Number of books: %v", count)
	}
	for i = 0; i < count; i++ {
		fmt.Printf("Book name: %v\nAuthor name: %v\nNumber of pages: %v", myLib[i].authorName, myLib[i].bookName, myLib[i].numberOfPages)
	}

}
