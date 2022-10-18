package section1

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Count the Arguments
//
//  Print the count of the command-line arguments
//
// INPUT
//  bilbo balbo bungo
//
// EXPECTED OUTPUT
//  There are 3 names.
// ---------------------------------------------------------

func Args1() {
	count := len(os.Args) - 1
	fmt.Printf("There are %d names.\n", count)
}
