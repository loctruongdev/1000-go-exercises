package masker

import (
	"fmt"
	"os"
)

func Masker() {
	queries := os.Args[1:]
	// fmt.Printf("len queries: %d\n", len(queries))

	if len(queries) == 0 {
		fmt.Println("Give me some things to search")
		return
	}

	var buf []byte

	for _, query := range queries {
		bquery := []byte(query)
		if string(bquery[:7]) == "http://" {

			for i := range bquery[7:] {
				bquery[7:][i] = '*'
			}
			// fmt.Println("catched http://")
			// fmt.Println("http://", string(bquery[7:]))
			// return
		}

		buf = append(buf, bquery...)
		buf = append(buf, ' ')
		// fmt.Println(string(buf))

	}
	// fmt.Printf("queries: %q\n", queries)
	fmt.Println(string(buf))

	/*
		1. Check whether there's a command line argument or not. If not, quit from the program with a message.
		done

		2. Create a byte buffer as big as the argument.

		3. Loop and detect the http:// patterns

			Copy the input character by character to the buffer

			If you detect http:// pattern, copy the http:// first, then copy the *s instead of the original link until you see whitespace character.

	*/
}
