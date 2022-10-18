package section3

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Fix
//
//  1. Uncomment the code
//
//  2. And fix the problems
//
//  3. BONUS: Simplify the code
// ---------------------------------------------------------

func FixCode() {
	names := [3]string{
		"Einstein",
		"Shepard",
		"Tesla",
	}

	// names := [...]string{
	// 	"Einstein",
	// 	"Shepard",
	// 	"Tesla",
	// }

	books := [5]string{
		"Kafka's Revenge",
		"Stay Golden",
	}

	fmt.Printf("%q\n", names)
	fmt.Printf("%q\n", books)
}
