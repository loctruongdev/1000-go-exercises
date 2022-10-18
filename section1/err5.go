package section1

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Days in a Month
//
//  Print the number of days in a given month.
//
// RESTRICTIONS
//  1. On a leap year, february should print 29. Otherwise, 28.
//
//     Set your computer clock to 2020 to see whether it works.
//
//  2. It should work case-insensitive. See below.
//
//     Search on Google: golang pkg strings ToLower
//
//  3. Get the current year using the time.Now()
//
//     Search on Google: golang pkg time now year
//
//
// EXPECTED OUTPUT
//
//  -----------------------------------------
//  Your solution should not accept invalid months
//  -----------------------------------------
//  go run main.go
//    Give me a month name
//
//  go run main.go sheep
//    "sheep" is not a month.
//
//  -----------------------------------------
//  Your solution should handle the leap years
//  -----------------------------------------
//  go run main.go january
//    "january" has 31 days.
//
//  go run main.go february
//    "february" has 28 days.
//
//  go run main.go march
//    "march" has 31 days.
//
//  go run main.go april
//    "april" has 30 days.
//
//  go run main.go may
//    "may" has 31 days.
//
//  go run main.go june
//    "june" has 30 days.
//
//  go run main.go july
//    "july" has 31 days.
//
//  go run main.go august
//    "august" has 31 days.
//
//  go run main.go september
//    "september" has 30 days.
//
//  go run main.go october
//    "october" has 31 days.
//
//  go run main.go november
//    "november" has 30 days.
//
//  go run main.go december
//    "december" has 31 days.
//
//  -----------------------------------------
//  Your solution should be case insensitive
//  -----------------------------------------
//  go run main.go DECEMBER
//    "DECEMBER" has 31 days.
//
//  go run main.go dEcEmBeR
//    "dEcEmBeR" has 31 days.
// ---------------------------------------------------------

func DaysInMonth() {

	if len(os.Args) != 2 {
		fmt.Println("Give me a month name")
		return
	}
	// get the current year and find out whether it's a leap year
	year := time.Now().Year()

	var feb int

	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		feb = 29
	} else {
		feb = 28
	}

	m := strings.ToLower(os.Args[1])

	if m == "january" {
		fmt.Printf("%q has 31 days.\n", os.Args[1])
	} else if m == "february" {
		fmt.Printf("%q has %d days.\n", os.Args[1], feb)
	} else if m == "march" {
		fmt.Printf("%q has 29 days.\n", os.Args[1])
	} else if m == "april" {
		fmt.Printf("%q has 30 days.\n", os.Args[1])
	} else if m == "may" {
		fmt.Printf("%q has 31 days.\n", os.Args[1])
	} else if m == "june" {
		fmt.Printf("%q has 30 days.\n", os.Args[1])
	} else if m == "july" {
		fmt.Printf("%q has 31 days.\n", os.Args[1])
	} else if m == "august" {
		fmt.Printf("%q has 31 days.\n", os.Args[1])
	} else if m == "september" {
		fmt.Printf("%q has 30 days.\n", os.Args[1])
	} else if m == "october" {
		fmt.Printf("%q has 31 days.\n", os.Args[1])
	} else if m == "november" {
		fmt.Printf("%q has 30 days.\n", os.Args[1])
	} else if m == "december" {
		fmt.Printf("%q has 31 days.\n", os.Args[1])
	} else {
		fmt.Printf("%q is not a month.\n", os.Args[1])
	}

}

func DaysInMonthb() {
	if len(os.Args) != 2 {
		fmt.Println("Give me a month name")
		return
	}

	// get the current year and find out whether it's a leap year
	year := time.Now().Year()
	leap := year%4 == 0 && (year%100 != 0 || year%400 == 0)

	// setting it to 28, saves me typing it below again
	days := 28

	month := os.Args[1]

	// case insensitive
	if m := strings.ToLower(month); m == "april" ||
		m == "june" ||
		m == "september" ||
		m == "november" {
		days = 30
	} else if m == "january" ||
		m == "march" ||
		m == "may" ||
		m == "july" ||
		m == "august" ||
		m == "october" ||
		m == "december" {
		days = 31
	} else if m == "february" {
		// try a "logical and operator" above.
		// like: `else if m == "february" && leap`
		if leap {
			days = 29
		}
	} else {
		fmt.Printf("%q is not a month.\n", month)
		return
	}

	fmt.Printf("%q has %d days.\n", month, days)
}
