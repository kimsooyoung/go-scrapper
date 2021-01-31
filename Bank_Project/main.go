package main

import (
	"fmt"

	"github.com/swimming/go-scrapper/Bank_Project/accounts"
	"github.com/swimming/go-scrapper/Bank_Project/mydict"
)

func ExceptionHandle(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	myAccount := accounts.NewAccount("Kim")
	myAccount.Deposit(100)
	fmt.Println(myAccount.Balance())
	// go doesn't have exception handling syntaxes like try - except
	// So such an expression will used often

	// got return, then check if there was an error or not
	err := myAccount.WithDraw(200)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(myAccount)
	fmt.Println("===========================")

	key := "name"
	value := "Kim"

	dictionary := mydict.Dictionary{}
	// fmt.Println(dictionary["name"])

	// Let's use this function with Type Method
	addErr := dictionary.Add(key, value) // '' <= char "" <= string
	ExceptionHandle(addErr)

	definition, _ := dictionary.Search(key)
	// ExceptionHandle(searchErr)
	fmt.Println(key, "+", definition)

	addErr2 := dictionary.Add(key, value)
	ExceptionHandle(addErr2)

	updateErr := dictionary.Update("Sex", "Male")
	ExceptionHandle(updateErr)

	deleteErr := dictionary.Delete("Location")
	ExceptionHandle(deleteErr)
}
