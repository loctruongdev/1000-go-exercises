package section1

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Multi Assign
//
//  1. Assign "go" to `lang` variable
//     and assign 2 to `version` variable
//     using a multiple assignment statement
//
//  2. Print the variables
//
// EXPECTED OUTPUT
//  go version 2
// ---------------------------------------------------------

func MultiAssign() {

	var (
		lang    string
		version int
	)
	// lang = "go"
	// version = 2
	lang, version = "go", 2

	fmt.Println(lang, "version", version)

}
