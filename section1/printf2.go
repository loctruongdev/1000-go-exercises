package section1

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Print Your Name and LastName
//
//  Print your name and lastname using Printf
//
// EXPECTED OUTPUT
//  My name is Inanc and my lastname is Gumus.
//
// BONUS
//  Store the formatting specifier (first argument of Printf)
//    in a variable.
//  Then pass it to printf
// ---------------------------------------------------------

func PrintNameLastname() {
	firstName, lastName := "Spring", "Truong"
	fmt.Printf("My name is %s and my last name is %s\n.", firstName, lastName)
}

func PrintNameLastnameb() {
	// fmt.Printf("My name is %s and my last name is %s\n.", "Spring", "Truong")

	// BONUS
	msg := "My name is %s and my last name is %s.\n"
	fmt.Printf(msg, "Spring", "Truong")
}
