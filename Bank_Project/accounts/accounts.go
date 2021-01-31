package accounts

// Account struct
// In Golang, Uppercase means public,
// you only can use struct in main.go if it starts with Uppercase
type Account struct {
	owner   string
	balance int
}

// Popular Pattern in go, make new struct with Fucntion, similar like Constructor
func NewAccount(owner_in string) *Account {
	account := Account{owner: owner_in, balance: 0}
	return &account
}

// Receiver makes function as a Struct Method
func (a Account) Deposit(amount int) {
	a.balance += amount
}

// Define Getter for Balance
func (a Account) Balance() int {
	return a.balance
}
