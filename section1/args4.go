package section1

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Greet More People
//
// RESTRICTIONS
//  1. Be sure to match the expected output below
//  2. Store the length of the arguments in a variable
//  3. Store all the arguments in variables as well
//
// INPUT
//  bilbo balbo bungo
//
// EXPECTED OUTPUT
//  There are 3 people!
//  Hello great bilbo !
//  Hello great balbo !
//  Hello great bungo !
//  Nice to meet you all.
// ---------------------------------------------------------

func Args4() {
	l := len(os.Args) - 1
	first := os.Args[1]
	second := os.Args[2]
	third := os.Args[3]
	fmt.Printf("There are %d people!\n", l)
	fmt.Printf("Hello great %s !\n", first)
	fmt.Printf("Hello great %s !\n", second)
	fmt.Printf("Hello great %s !\n", third)
	fmt.Println("Nice to meet you all.")

}

func Args4b() {
	var (
		l  = len(os.Args) - 1
		n1 = os.Args[1]
		n2 = os.Args[2]
		n3 = os.Args[3]
	)

	fmt.Println("There are", l, "people!")
	fmt.Println("Hello great", n1, "!")
	fmt.Println("Hello great", n2, "!")
	fmt.Println("Hello great", n3, "!")
	fmt.Println("Nice to meet you all.")

}
