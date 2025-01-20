package main

import (
	"fmt"
)

type Book struct {
	Title     string
	Author    string
	Pages     int
	Available bool
}

// Create a new book and return it
func AddBook(n string, a string, p int, av bool) Book {
	return Book{
		Title:     n,
		Author:    a,
		Pages:     p,
		Available: av,
	}
}

// Borrow a book (set Available to false)
func (b *Book) BorrowBook() {
	b.Available = false
}

// Return a book (set Available to true)
func (b *Book) ReturnBook() {
	b.Available = true
}

func main() {
	// Initialize a collection of books
	var library []Book

	// Add books to the library
	book1 := AddBook("Go Programming", "John Doe", 300, true)
	book2 := AddBook("Advanced Go", "Jane Doe", 500, true)

	library = append(library, book1, book2)

	// Print the library
	fmt.Println("Library Collection:")
	for _, b := range library {
		fmt.Printf("Title: %s, Author: %s, Pages: %d, Available: %t\n", b.Title, b.Author, b.Pages, b.Available)
	}

	// Borrow the first book
	fmt.Println("\nBorrowing the first book...")
	library[0].BorrowBook()
	fmt.Printf("Title: %s, Available: %t\n", library[0].Title, library[0].Available)

	// Return the first book
	fmt.Println("\nReturning the first book...")
	library[0].ReturnBook()
	fmt.Printf("Title: %s, Available: %t\n", library[0].Title, library[0].Available)
}
