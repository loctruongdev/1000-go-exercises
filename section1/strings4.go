package section1

import (
	"fmt"
	"os"
	"unicode/utf8"
)

// ---------------------------------------------------------
// EXERCISE: Count the Chars
//
//  1. Change the following program to work with unicode
//     characters.
//
// INPUT
//  "İNANÇ"
//
// EXPECTED OUTPUT
//  5
// ---------------------------------------------------------

func CountChars() {
	length := utf8.RuneCountInString(os.Args[1])
	fmt.Println(length)
}
