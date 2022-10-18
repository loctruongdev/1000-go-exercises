package section1

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Arg Count
//
//  1. Get arguments from command-line.
//  2. Print the expected outputs below depending on the number
//     of arguments.
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me args
//
//  go run main.go hello
//    There is one: "hello"
//
//  go run main.go hi there
//    There are two: "hi there"
//
//  go run main.go I wanna be a gopher
//    There are 5 arguments
// ---------------------------------------------------------

func ArgCount() {
	l := len(os.Args) - 1
	if l == 0 {
		fmt.Println("Give me args")
	} else {
		var msg string
		for i := 1; i <= l; i++ {
			msg = msg + os.Args[i] + ` `
		}
		fmt.Printf("There is %d argument(s): %s\n", l, msg)
	}

}

func ArgCountb() {

	args := []string{"minh", "loc"}
	var msg string
	for i := 0; i <= len(args)-1; i++ {
		msg = msg + " " + args[i]
	}
	fmt.Printf("There is argument(s): %s\n", msg)

}

func ArgCountc() {
	var (
		args = os.Args
		l    = len(args) - 1
	)

	if l == 0 {
		fmt.Println("Give me args")
	} else if l == 1 {
		fmt.Printf("There is one: %q\n", args[1])
	} else if l == 2 {
		fmt.Printf(
			`There are two: "%s %s"`+"\n",
			args[1], args[2],
		)
	} else {
		fmt.Printf("There are %d arguments\n", l)
	}
}
