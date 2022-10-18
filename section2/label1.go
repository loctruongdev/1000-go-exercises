package section2

import (
	"fmt"
	"os"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Case Insensitive Search
//
//  Allow for case-insensitive searching
//
// EXAMPLE
//  Let's say that the user runs the program like this:
//    go run main.go LAZY
//
//  Or like this: go run main.go lAzY
//  Or like this: go run main.go lazy
//
//  For all cases above, the program should find
//  the "lazy" keyword.
// ---------------------------------------------------------

const corpus2 = "lazy cat jumps again and again and again"

func CaseInsensitiveSearch() {
	words := strings.Fields(corpus2)
	query := os.Args[1:]

queries:
	for _, q := range query {
		// case insensitive search
		q = strings.ToLower(q)

	search:
		for i, w := range words {
			switch q {
			case "and", "or", "the": // filter
				break search
			}

			if q == w {
				fmt.Printf("#%-2d: %q\n", i+1, w)
				continue queries
			}
		}
	}
}
