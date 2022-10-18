package section1

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Convert
//
//  Convert the if statement to a switch statement.
// ---------------------------------------------------------

const (
	usagex        = "Usage: [username] [password]"
	errUserx      = "Access denied for %q.\n"
	errPwdx       = "Invalid password for %q.\n"
	accessOKx     = "Access granted to %q.\n"
	userx, userx2 = "jack", "inanc"
	passx, passx2 = "1888", "1879"
)

func ConvertIf() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println(usagex)
		return
	}

	u, p := args[1], args[2]

	//
	// REFACTOR THIS TO A SWITCH
	//
	if u != userx && u != userx2 {
		fmt.Printf(errUserx, u)
	} else if u == userx && p == passx {
		fmt.Printf(accessOKx, u)
	} else if u == userx2 && p == passx2 {
		fmt.Printf(accessOKx, u)
	} else {
		fmt.Printf(errPwdx, u)
	}
}

func ConvertSwitch() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println(usagex)
		return
	}

	u, p := args[1], args[2]

	//
	// REFACTOR THIS TO A SWITCH
	//
	switch {
	case u != userx && u != userx2:
		fmt.Printf(errUserx, u)
	case u == userx && p == pass:
		fallthrough

	case u == userx2 && p == passx2:
		fmt.Printf(accessOKx, u)
	default:
		fmt.Printf(errPwdx, u)
	}

	// if u != user && u != user2 {
	// 	fmt.Printf(errUser, u)
	// } else if u == user && p == pass {
	// 	fmt.Printf(accessOK, u)
	// } else if u == user2 && p == pass2 {
	// 	fmt.Printf(accessOK, u)
	// } else {
	// 	fmt.Printf(errPwd, u)
	// }
}
