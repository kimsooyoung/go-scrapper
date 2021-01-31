package main

import (
	"fmt"
	"strings"
)

// you must specify the type of variables in function
// also needs to specify return types too.
func multiply(a int, b int) int {
	return a * b
}

// function with multiple returns
func lenAndUpper(name string) (int, string) {
	// defer runs right after this function done
	defer fmt.Println("lenAndUpper done... ")
	return len(name), strings.ToUpper(name)
}

func superAdd(numbers ...int) int {
	total := 0

	for _, val := range numbers {
		total += val
		fmt.Print(val, " ")
	}
	fmt.Println()

	return total
}

// naked return function
func nakedLenAndUpper(name string) (length int, upperName string) {
	length = len(name)
	upperName = strings.ToUpper(name)
	return
}

// function with multiple arguments
func multipleArgs(words ...string) {
	fmt.Println(words)
}

// if-else
// variable expression
func canIDrink(age int) bool {

	// you can assign variable in if expression
	if koreanAge := age + 2; koreanAge > 18 {
		return true
	}
	return false
}

func canIDrinkSwitch(age int) bool {

	// switch expression, you can also use variable expression in here too.
	switch koreanAge := age + 2; {
	case koreanAge < 18:
		return false
	case koreanAge == 18:
		return true
	case koreanAge > 20:
		return true
	}
}

func main() {
	const name string = "Swimming"
	// name = "Changed" => invalid, const variable

	var name2 string = "Jelly"
	same_name := "Swimming" // same with this => var same_name string = "Swimming"
	name2 = "Belly"         // => possible

	fmt.Println(name2)
	fmt.Println(same_name)

	fmt.Println(multiply(2, 2))

	const friend_name string = "James"

	// you can ignore return values with underscore _
	nameLen, upperName := lenAndUpper(friend_name)
	fmt.Println(nameLen, upperName)

	multipleArgs("This ", "is ", "test ", "sentences")

	summation := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(summation)

	fmt.Println(canIDrink(15))
}
