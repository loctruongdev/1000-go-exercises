package section1

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Richter Scale #2
//
//  Repeat the previous exercise.
//
//  But, this time, get the description and print the
//  corresponding richter scale.
//
//  See the expected outputs.
//
// ---------------------------------------------------------
// MAGNITUDE                    DESCRIPTION
// ---------------------------------------------------------
// Less than 2.0                micro
// 2.0 to less than 3.0         very minor
// 3.0 to less than 4.0         minor
// 4.0 to less than 5.0         light
// 5.0 to less than 6.0         moderate
// 6.0 to less than 7.0         strong
// 7.0 to less than 8.0         major
// 8.0 to less than 10.0        great
// 10.0 or more                 massive
//
// EXPECTED OUTPUT
//  go run main.go
//   Tell me the magnitude of the earthquake in human terms.
//
//  go run main.go alien
//   alien's richter scale is unknown
//
//  go run main.go micro
//   micro's richter scale is less than 2.0
//
//  go run main.go "very minor"
//   very minor's richter scale is 2 - 2.9
//
//  go run main.go minor
//   minor's richter scale is 3 - 3.9
//
//  go run main.go light
//   light's richter scale is 4 - 4.9
//
//  go run main.go moderate
//   moderate's richter scale is 5 - 5.9
//
//  go run main.go strong
//   strong's richter scale is 6 - 6.9
//
//  go run main.go major
//   major's richter scale is 7 - 7.9
//
//  go run main.go great
//   great's richter scale is 8 - 9.9
//
//  go run main.go massive
//   massive's richter scale is 10+
// ---------------------------------------------------------

func RichterScales2() {

	l := len(os.Args)
	var r string

	// if l == 2 {
	// 	r = os.Args[1]
	// } else if l == 3 {
	// 	r = os.Args[1] + ` ` + os.Args[2]
	// } else if l == 1 {
	// 	fmt.Println("Tell me the magnitude of the earthquake in human terms.")
	// 	r = "main"
	// }

	switch l {
	case 1:
		{
			fmt.Println("Tell me the magnitude of the earthquake in human terms.")
			r = "main"
		}
	case 2:
		{
			r = os.Args[1]
		}
	case 3:
		{
			r = os.Args[1] + ` ` + os.Args[2]
		}
	}

	var scale string

	switch r {

	case "micro":
		{
			scale = "less than 2.0"
		}
	case "very minor":
		{
			scale = "2 - 2.9"
		}
	case "minor":
		{
			scale = "3 - 3.9"
		}
	case "light":
		{
			scale = "4 - 4.9"
		}
	case "moderate":
		{
			scale = "5 - 5.9"
		}
	case "strong":
		{
			scale = "6 - 6.9"
		}
	case "major":
		{
			scale = "7 - 7.9"
		}
	case "great":
		{
			scale = "8 - 8.9"
		}
	case "massive":
		{
			scale = "10+"
		}
	default:
		scale = "unknown"

	}
	if r != "main" {
		fmt.Printf("%s's richter scale is %s\n", r, scale)
	}

}

func RichterScales2b() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Tell me the magnitude of the earthquake in human terms.")
		return
	}

	var richter string

	desc := args[1]

	switch desc {
	case "micro":
		richter = "less than 2.0"
	case "very minor":
		richter = "2 - 2.9"
	case "minor":
		richter = "3 - 3.9"
	case "light":
		richter = "4 - 4.9"
	case "moderate":
		richter = "5 - 5.9"
	case "strong":
		richter = "6 - 6.9"
	case "major":
		richter = "7 - 7.9"
	case "great":
		richter = "8 - 9.9"
	case "massive":
		richter = "10+"
	default:
		richter = "unknown"
	}

	fmt.Printf(
		"%s's richter scale is %s\n",
		desc, richter,
	)
}
