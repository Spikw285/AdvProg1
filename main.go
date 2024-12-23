package main

import (
	"Assignment1/Bank"
	"Assignment1/Employees"
	"Assignment1/Library"
	"Assignment1/Shapes"
	"fmt"
)

func main() {
	// Bank example
	bankAccount := &Bank.BankAccount{
		AccountNumber: "123456789",
		HolderName:    "John Doe",
		Balance:       1000.00,
	}

	// Performing bank transactions
	transactions := []float64{500, -200, 1000}
	Bank.Transaction(bankAccount, transactions)
	bankAccount.GetBalance()

	// Employees example
	fullTimeEmployee := Employees.FullTimeEmployee{
		ID:     1,
		Name:   "Alice",
		Salary: 50000,
	}

	partTimeEmployee := Employees.PartTimeEmployee{
		ID:          2,
		Name:        "Bob",
		HourlyRate:  20,
		HoursWorked: 15,
	}

	company := Employees.Company{
		Employees: make(map[string]Employees.Employee),
	}

	company.AddEmployee(fullTimeEmployee)
	company.AddEmployee(partTimeEmployee)

	// Listing all employees
	company.ListEmployees()

	// Library example
	library := Library.NewLibrary()

	book1 := Library.Book{
		ID:         "1",
		Title:      "The Go Programming Language",
		Author:     "Alan Donovan",
		IsBorrowed: false,
	}

	book2 := Library.Book{
		ID:         "2",
		Title:      "Clean Code",
		Author:     "Robert C. Martin",
		IsBorrowed: false,
	}

	library.AddBook(book1)
	library.AddBook(book2)

	// Borrow and return books
	library.BorrowBook("1")
	library.ListBooks()
	library.ReturnBook("1")
	library.ListBooks()

	// Shapes example
	rectangle := Shapes.Rectangle{Length: 5, Width: 3}
	circle := Shapes.Circle{Radius: 7}
	square := Shapes.Square{Length: 4}
	triangle := Shapes.Triangle{SideA: 3, SideB: 4, SideC: 5}

	// Displaying area and perimeter for shapes
	fmt.Printf("Rectangle - Area: %.2f, Perimeter: %.2f\n", rectangle.Area(), rectangle.Perimeter())
	fmt.Printf("Circle - Area: %.2f, Perimeter: %.2f\n", circle.Area(), circle.Perimeter())
	fmt.Printf("Square - Area: %.2f, Perimeter: %.2f\n", square.Area(), square.Perimeter())
	fmt.Printf("Triangle - Area: %.2f, Perimeter: %.2f\n", triangle.Area(), triangle.Perimeter())
}
