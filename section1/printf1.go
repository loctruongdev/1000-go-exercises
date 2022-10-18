package section1

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Print Your Age
//
//  Print your age using Printf
//
// EXPECTED OUTPUT
//  I'm 30 years old.
//
// NOTE
//  You should change 30 to your age, of course.
// ---------------------------------------------------------

func PrintYourAge() {
	age := 30
	fmt.Printf("I'm %d years old\n", age)
}

func MyPrintf1b() {
	fmt.Printf("I'm %d years old\n", 35)
}
