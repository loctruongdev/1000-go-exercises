package section1

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// ---------------------------------------------------------
// EXERCISE: Right Trim It
//
//  1. Look at the documentation of strings package
//  2. Find a function that trims the spaces from
//     only the right-most part of the given string
//  3. Trim it from the right part only
//  4. Print the number of characters it contains.
//
// RESTRICTION
//  Your program should work with unicode string values.
//
// EXPECTED OUTPUT
//  5
// ---------------------------------------------------------

func RightTrimIt() {
	// currently it prints 19
	// it should print 7

	name := "  inanç           "
	fmt.Println(len(name))
	fmt.Println(utf8.RuneCountInString(strings.TrimRight(name, " ")))
}
