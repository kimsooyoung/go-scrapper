package main

import (
	"fmt"

	"github.com/swimming/go-scrapper/Bank_Project/accounts"
)

func main() {
	myAccount := accounts.NewAccount("Kim")
	myAccount.Deposit(100)
	fmt.Println(myAccount.Balance())
}
