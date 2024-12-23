package Library

import "fmt"

type Book struct {
	ID         string
	Title      string
	Author     string
	IsBorrowed bool
}

type Library struct {
	Books map[string]Book
}

func NewLibrary() *Library {
	fmt.Println("Library initialized.")
	return &Library{
		Books: make(map[string]Book),
	}
}

func (lib *Library) AddBook(book Book) {
	if _, exists := lib.Books[book.ID]; exists {
		fmt.Println("Error: Book with this ID already exists in the library.")
		return
	}
	lib.Books[book.ID] = book
	fmt.Println(book.Title, "was added to the library.")
}

func (lib *Library) BorrowBook(id string) {
	if book, exists := lib.Books[id]; exists {
		if book.IsBorrowed {
			fmt.Printf("Error: Book '%s' is already borrowed.\n", book.Title)
		} else {
			book.IsBorrowed = true
			lib.Books[id] = book
			fmt.Printf("Book '%s' has been borrowed.\n", book.Title)
		}
	} else {
		fmt.Printf("Error: Book with ID '%s' was not found.\n", id)
	}
}

func (lib *Library) ReturnBook(id string) {
	if book, exists := lib.Books[id]; exists {
		if !book.IsBorrowed {
			fmt.Printf("Error: Book '%s' is already available in the library.\n", book.Title)
		} else {
			book.IsBorrowed = false
			lib.Books[id] = book
			fmt.Printf("Book '%s' has been successfully returned.\n", book.Title)
		}
	} else {
		fmt.Printf("Error: Book with ID '%s' was not found.\n", id)
	}
}

func (lib *Library) ListBooks() {
	fmt.Println("Books in the library:")
	if len(lib.Books) == 0 {
		fmt.Println("The library is empty.")
		return
	}
	for _, book := range lib.Books {
		status := "Available"
		if book.IsBorrowed {
			status = "Borrowed"
		}
		fmt.Printf("- ID: %s | Title: %s | Author: %s | Status: %s\n", book.ID, book.Title, book.Author, status)
	}
}
