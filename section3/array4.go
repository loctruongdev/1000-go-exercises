package section3

import (
	"fmt"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Refactor to Ellipsis
//
//  1. Use the 03-array-literal exercise
//
//  2. Refactor the length of the array literals to ellipsis
//
//    This means: Use the ellipsis instead of defining the array's length
//                manually.
//
// EXPECTED OUTPUT
//   The output should be the same as the 03-array-literal exercise.
// ---------------------------------------------------------

func RefactorEllipsis() {
	names := [...]string{"Viggo", "Radis", "Bean"}

	distances := [...]int{
		1000,
		800,
		100,
		600,
		500,
	}

	data := [...]byte{97, 98, 99, 100, 101}

	ratios := [...]float64{23.495}

	alives := [...]bool{true, false, false, true}

	zero := [...]byte{}

	separator := "\n" + strings.Repeat("=", 20) + "\n"

	fmt.Print("names", separator)
	for i := 0; i < len(names); i++ {
		fmt.Printf("names[%d]: %s\n", i, names[i])
	}

	fmt.Print("\ndistances", separator)
	for i := 0; i < len(distances); i++ {
		fmt.Printf("distances[%d]: %d\n", i, distances[i])
	}

	fmt.Print("\ndata", separator)
	for i := 0; i < len(data); i++ {
		fmt.Printf("data[%d]: %d\n", i, data[i])
	}

	fmt.Print("\nratios", separator)
	for i := 0; i < len(ratios); i++ {
		fmt.Printf("ratios[%d]: %.2f\n", i, ratios[i])
	}

	fmt.Print("\nalives", separator)
	for i := 0; i < len(alives); i++ {
		fmt.Printf("alives[%d]: %t\n", i, alives[i])
	}

	fmt.Print("\nzero", separator)
	for i := 0; i < len(zero); i++ {
		fmt.Printf("data[%d]: %d\n", i, zero[i])
	}

	fmt.Printf(`
	%s
	FOR RANGES
	%[1]s
	`, strings.Repeat("~", 30))

	fmt.Print("names", separator)
	for i, v := range names {
		fmt.Printf("names[%d]: %s\n", i, v)
	}

	fmt.Print("\ndistances", separator)
	for i, v := range distances {
		fmt.Printf("distances[%d]: %d\n", i, v)
	}

	fmt.Print("\ndata", separator)
	for i, v := range data {
		fmt.Printf("data[%d]: %d\n", i, v)
	}

	fmt.Print("\nratios", separator)
	for i, v := range ratios {
		fmt.Printf("ratios[%d]: %.2f\n", i, v)
	}

	fmt.Print("\nalives", separator)
	for i, v := range alives {
		fmt.Printf("alives[%d]: %t\n", i, v)
	}

	fmt.Print("\nzero", separator)
	for i, v := range zero {
		fmt.Printf("names[%d]: %d\n", i, v)
	}
}
