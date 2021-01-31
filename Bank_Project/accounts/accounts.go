package accounts

import (
	"errors"
	"fmt"
)

// Account struct
// In Golang, Uppercase means public,
// you only can use struct in main.go if it starts with Uppercase
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("There's no money")

// Popular Pattern in go, make new struct with Fucntion, similar like Constructor
func NewAccount(owner_in string) *Account {
	account := Account{owner: owner_in, balance: 0}
	return &account
}

// Receiver makes function as a Struct Method
// Pointer receiver used in here a *Account
// this
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Define Getter for Balance
func (a Account) Balance() int {
	return a.balance
}

func (a Account) Owner() string {
	return a.owner
}

func (a *Account) WithDraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

// String() is an special function and works like __str__ in python class
func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account. \nBalance : ", a.Balance())
}
