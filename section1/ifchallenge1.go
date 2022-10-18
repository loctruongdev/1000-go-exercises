package section1

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// CHALLENGE #1
//  Create a user/password protected program.
//
// EXAMPLE USER
//  username: jack
//  password: 1888
//
// EXPECTED OUTPUT
//  go run main.go
//    Usage: [username] [password]
//
//  go run main.go albert
//    Usage: [username] [password]
//
//  go run main.go hacker 42
//    Access denied for "hacker".
//
//  go run main.go jack 6475
//    Invalid password for "jack".
//
//  go run main.go jack 1888
//    Access granted to "jack".
// ---------------------------------------------------------

func Challenge1() {
	if len(os.Args) == 3 {

		u, p := os.Args[1], os.Args[2]
		// password, _ := strconv.Atoi(os.Args[2])

		if u == "jack" && p == "1888" {
			fmt.Printf("Access granted to %q.\n", u)
		} else if u == "jack" && p != "1888" {
			fmt.Printf("Invalid password for %q.\n", u)
		} else {
			fmt.Printf("Access denied for %q.\n", u)
		}
	} else {
		fmt.Println("Usage: [username] [password]")
		return
	}
}

func Challenge1b() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println("Usage: [username] [password]")
		return
	}

	u, p := args[1], args[2]

	if u != "jack" {
		fmt.Printf("Access denied for %q.\n", u)
	} else if p != "1888" {
		fmt.Printf("Invalid password for %q.\n", u)
	} else {
		fmt.Printf("Access granted to %q.\n", u)
	}
}
