package main

import (
	"Assignment1/Library"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	lib := Library.NewLibrary()
	scanner := bufio.NewScanner(os.Stdin)

	idPattern := regexp.MustCompile(`^\d+$`)
	wordPattern := regexp.MustCompile(`^[a-zA-Z]+$`)

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Add - Add the book")
		fmt.Println("2. Borrow - Borrow the book")
		fmt.Println("3. Return - Return the book")
		fmt.Println("4. List - Show all books")
		fmt.Println("5. Exit - Exit program")

		fmt.Print("Choose action: ")
		scanner.Scan()
		command := strings.ToLower(strings.TrimSpace(scanner.Text()))

		switch command {
		case "1", "add":
			fmt.Print("Enter book ID (only digits): ")
			scanner.Scan()
			id := strings.TrimSpace(scanner.Text())
			if !idPattern.MatchString(id) {
				fmt.Println("Error: Book ID must contain only digits.")
				continue
			}

			fmt.Print("Enter book title (at least 4 characters or 1 word): ")
			scanner.Scan()
			title := strings.TrimSpace(scanner.Text())
			if len(title) < 4 && !strings.Contains(title, " ") {
				fmt.Println("Error: Book title must be at least 4 characters or contain 1 word.")
				continue
			}

			fmt.Print("Enter book author (at least 3 characters): ")
			scanner.Scan()
			author := strings.TrimSpace(scanner.Text())
			if len(author) < 3 || !wordPattern.MatchString(author) {
				fmt.Println("Error: Author's name must be at least 3 characters.")
				continue
			}

			book := Library.Book{
				ID:     id,
				Title:  title,
				Author: author,
			}
			lib.AddBook(book)

		case "2", "borrow":
			fmt.Print("Enter book ID for borrowing: ")
			scanner.Scan()
			id := strings.TrimSpace(scanner.Text())
			if !idPattern.MatchString(id) {
				fmt.Println("Error: Book ID must contain only digits.")
				continue
			}
			lib.BorrowBook(id)

		case "3", "return":
			fmt.Print("Enter book ID for returning: ")
			scanner.Scan()
			id := strings.TrimSpace(scanner.Text())
			if !idPattern.MatchString(id) {
				fmt.Println("Error: Book ID must contain only digits.")
				continue
			}
			lib.ReturnBook(id)

		case "4", "list":
			lib.ListBooks()

		case "5", "exit":
			fmt.Println("Exiting program...")
			return

		default:
			fmt.Println("Unknown command. Please choose a valid option from the menu.")
		}
	}
}
