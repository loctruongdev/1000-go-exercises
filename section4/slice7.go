package section4

import (
	"bytes"
	"fmt"
)

// ---------------------------------------------------------
// EXERCISE: Append
//
//  Please follow the instructions within the code below.
//
// EXPECTED OUTPUT
//  They are equal.
//
// HINTS
//  bytes.Equal function allows you to compare two byte
//  slices easily. Check its documentation: go doc bytes.Equal
// ---------------------------------------------------------

func Append() {
	// 1. uncomment the code below
	// png, header := []byte{'P', 'N', 'G'}, []byte{}

	// 2. append elements to header to make it equal with the png slice

	// 3. compare the slices using the bytes.Equal function

	// 4. print whether they're equal or not

	png, header := []byte{'P', 'N', 'G'}, []byte{}

	// header = append(header, 'P', 'N', 'G')
	header = append(header, png...)

	if bytes.Equal(png, header) {
		fmt.Println("They are equal.")
	} else {
		fmt.Println("They are NOT equal")
	}
}
