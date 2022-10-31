package section11

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Uppercaser
//
//  Use a scanner to convert the lines to uppercase, and
//  print them.
//
//  1. Feed the shakespeare.txt to your program.
//
//  2. Scan the input using a new Scanner.
//
//  3. Print each line.
//
// EXPECTED OUTPUT
//  Please run the solution to see the expected output.
// ---------------------------------------------------------

func Uppercaser() {
	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		content := strings.ToUpper(in.Text())
		fmt.Println(content)
	}

}
