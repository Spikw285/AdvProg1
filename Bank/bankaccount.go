package Bank

import "fmt"

// Define the BankAccount struct
type BankAccount struct {
	AccountNumber string
	HolderName    string
	Balance       float64
}

// Implement the Deposit method for BankAccount
func (b *BankAccount) Deposit(amount float64) {
	// Add the amount to the balance
	b.Balance += amount
	fmt.Printf("Deposited %.2f. New balance: %.2f\n", amount, b.Balance)
}

// Implement the Withdraw method for BankAccount
func (b *BankAccount) Withdraw(amount float64) {
	// Deduct the amount from the balance only if sufficient balance exists
	if b.Balance >= amount {
		b.Balance -= amount
		fmt.Printf("Withdrew %.2f. New balance: %.2f\n", amount, b.Balance)
	} else {
		fmt.Println("Insufficient funds for withdrawal.")
	}
}

// Implement the GetBalance method for BankAccount
func (b *BankAccount) GetBalance() {
	// Print the current balance
	fmt.Printf("Current balance: %.2f\n", b.Balance)
}

// Function to process a slice of transactions
func Transaction(account *BankAccount, transactions []float64) {
	// Iterate over the transaction slice
	for _, transaction := range transactions {
		if transaction > 0 {
			// Deposit if transaction is positive
			account.Deposit(transaction)
		} else if transaction < 0 {
			// Withdraw if transaction is negative
			account.Withdraw(-transaction)
		}
	}
}
