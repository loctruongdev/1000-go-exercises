package exercises

import (
	"fmt"
	"os"
	"path"
	"strings"
	"unicode/utf8"
)

////// VARIABLE EXERCISES /////

// ---------------------------------------------------------
// EXERCISE: Make It Blue
//
//  1. Change `color` variable's value to "blue"
//
//  2. Print it
//
// EXPECTED OUTPUT
//  blue
// ---------------------------------------------------------

func Variables1() {

	color := "green"

	color = "blue"

	fmt.Println(color)
}

// ---------------------------------------------------------
// EXERCISE: Variables To Variables
//
//  1. Change the value of `color` variable to "dark green"
//
//  2. Do not assign "dark green" to `color` directly
//
//     Instead, use the `color` variable again
//     while assigning to `color`
//
//  3. Print it
//
// RESTRICTIONS
//  WRONG ANSWER, DO NOT DO THIS:
//  `color = "dark green"`
//
// HINT
//  + operator can concatenate string values
//
//  Example:
//
//  "h" + "e" + "y" returns "hey"
//
// EXPECTED OUTPUT
//  dark green
// ---------------------------------------------------------

func Variables2() {

	color := "green"

	color = "dark" + " " + color

	fmt.Println(color)
}

// ---------------------------------------------------------
// EXERCISE: Assign With Expressions
//
//  1. Multiply 3.14 with 2 and assign it to `n` variable
//
//  2. Print the `n` variable
//
// HINT
//  Example: 3 * 2 = 6
//  * is the product operator (it multiplies numbers)
//
// EXPECTED OUTPUT
//  6.28
// ---------------------------------------------------------

func Variables3() {
	n := 0.
	n = 3.14 * 2

	fmt.Println(n)
}

// ---------------------------------------------------------
// EXERCISE: Find the Rectangle's Perimeter
//
//  1. Find the perimeter of a rectangle
//     Its width is  5
//     Its height is 6
//
//  2. Assign the result to the `perimeter` variable
//
//  3. Print the `perimeter` variable
//
// HINT
//  Formula = 2 times the width and height
//
// EXPECTED OUTPUT
//  22
//
// BONUS
//  Find more formulas here and calculate them in new programs
//  https://www.mathsisfun.com/area.html
// ---------------------------------------------------------

func Variables4a() {
	var (
		perimeter     int
		width, height = 5, 6
	)

	perimeter = (width + height) * 2

	fmt.Println(perimeter)
}

func Variables4b(w, h int) {
	perimeter := (w + h) * 2

	fmt.Println(perimeter)
}

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

func Variables5() {

	var (
		lang    string
		version int
	)
	// lang = "go"
	// version = 2
	lang, version = "go", 2

	fmt.Println(lang, "version", version)

}

// ---------------------------------------------------------
// EXERCISE: Multi Assign #2
//
//  1. Assign the correct values to the variables
//     to match to the EXPECTED OUTPUT below
//
//  2. Print the variables
//
// HINT
//  Use multiple Println calls to print each sentence.
//
// EXPECTED OUTPUT
//  Air is good on Mars
//  It's true
//  It is 19.5 degrees
// ---------------------------------------------------------

func Variables6() {
	var (
		planet string
		isTrue bool
		temp   float64
	)

	planet, isTrue, temp = "Mars", true, 19.5
	fmt.Println("Air is good on ", planet)
	fmt.Println("It's ", isTrue)
	//fmt.Printf("It is %v degrees", temp)
	fmt.Println("It is", temp, "degrees")
}

// ---------------------------------------------------------
// EXERCISE: Multi Short Func
//
// 	1. Declare two variables using multiple short declaration syntax
//
//  2. Initialize the variables using `multi` function below
//
//  3. Discard the 1st variable's value in the declaration
//
//  4. Print only the 2nd variable
//
// NOTE
//  You should use `multi` function
//  while initializing the variables
//
// EXPECTED OUTPUT
//  4
// ---------------------------------------------------------

func Variables7() {
	_, b := multi()

	fmt.Println(b)
}

func multi() (int, int) {
	return 5, 4
}

// ---------------------------------------------------------
// EXERCISE: Swapper
//
//  1. Change `color` to "orange"
//     and `color2` to "green" at the same time
//
//     (use multiple-assignment)
//
//  2. Print the variables
//
// EXPECTED OUTPUT
//  orange green
// ---------------------------------------------------------
func Variables8() {
	color, color2 := "red", "blue"
	color, color2 = "orange", "green"
	fmt.Println(color, color2)
}

// ---------------------------------------------------------
// EXERCISE: Swapper #2
//
//  1. Swap the values of `red` and `blue` variables
//
//  2. Print them
//
// EXPECTED OUTPUT
//  blue red
// ---------------------------------------------------------

func Variables9() {
	color, color2 := "red", "blue"
	color, color2 = color2, color
	fmt.Println(color, color2)
}

// ---------------------------------------------------------
// EXERCISE: Discard The File
//
//  1. Print only the directory using `path.Split`
//
//  2. Discard the file part
//
// RESTRICTION
//  Use short declaration
//
// EXPECTED OUTPUT
//  secret/
// ---------------------------------------------------------

func Variables10() {
	dir, _ := path.Split("secret/file.txt")
	fmt.Println(dir)
}

////// TYPE-CONVERSION EXERCISES /////
// ---------------------------------------------------------
// EXERCISE: Convert and Fix
//
//  Fix the code by using the conversion expression.
//
// EXPECTED OUTPUT
//  15.5
// ---------------------------------------------------------

func TypeC1() {
	a, b := 10, 5.5
	fmt.Println(float64(a) + b)
}

// ---------------------------------------------------------
// EXERCISE: Convert and Fix #2
//
//  Fix the code by using the conversion expression.
//
// EXPECTED OUTPUT
//  10.5
// ---------------------------------------------------------

func TypeC2() {
	a, b := 10, 5.5
	a = int(b)
	fmt.Println(float64(a) + b)
}

// ---------------------------------------------------------
// EXERCISE: Convert and Fix #3
//
//  Fix the code.
//
// EXPECTED OUTPUT
//  5.5
// ---------------------------------------------------------

func TypeC3() {
	fmt.Println(5.5)
}

// ---------------------------------------------------------
// EXERCISE: Convert and Fix #4
//
//  Fix the code.
//
// EXPECTED OUTPUT
//  9.5
// ---------------------------------------------------------

func TypeC4() {
	age := 2
	fmt.Println(7.5 + float64(age))
}

// ---------------------------------------------------------
// EXERCISE: Convert and Fix #5
//
//  Fix the code.
//
// HINTS
//   maximum of int8  can be 127
//   maximum of int16 can be 32767
//
// EXPECTED OUTPUT
//  1127
// ---------------------------------------------------------

func TypeC5() {
	// DO NOT TOUCH THESE VARIABLES
	min := int8(127)
	max := int16(1000)

	// FIX THE CODE HERE
	fmt.Println(max + int16(min))
}

/* EXPLANATION
`int8(max)` destroys the information of max
It reduces it to 127
Which is the maximum value of int8

Correct conversion is int16(min)
Because, int16 > int8
When you do so, min doesn't lose information
*/

////// COMMAND LINE ARGUMENTS EXERCISES /////

// ---------------------------------------------------------
// EXERCISE: Count the Arguments
//
//  Print the count of the command-line arguments
//
// INPUT
//  bilbo balbo bungo
//
// EXPECTED OUTPUT
//  There are 3 names.
// ---------------------------------------------------------

func ClArgs1() {
	count := len(os.Args) - 1
	fmt.Printf("There are %d names.\n", count)
}

// ---------------------------------------------------------
// EXERCISE: Print the Path
//
//  Print the path of the running program by getting it
//  from `os.Args` variable.
//
// HINT
//  Use `go build` to build your program.
//  Then run it using the compiled executable program file.
//
// EXPECTED OUTPUT SHOULD INCLUDE THIS
//  myprogram
// ---------------------------------------------------------

func ClArgs2() {
	fmt.Println(os.Args[0])
}

// ---------------------------------------------------------
// EXERCISE: Print Your Name
//
//  Get it from the first command-line argument
//
// INPUT
//  Call the program using your name
//
// EXPECTED OUTPUT
//  It should print your name
//
// EXAMPLE
//  go run main.go inanc
//
//    inanc
//
// BONUS: Make the output like this:
//
//  go run main.go inanc
//    Hi inanc
//    How are you?
// ---------------------------------------------------------

func ClArgs3() {
	// fmt.Println(os.Args[1])

	fmt.Println("Hi", os.Args[1])
	fmt.Println("How are you?")

}

// ---------------------------------------------------------
// EXERCISE: Greet More People
//
// RESTRICTIONS
//  1. Be sure to match the expected output below
//  2. Store the length of the arguments in a variable
//  3. Store all the arguments in variables as well
//
// INPUT
//  bilbo balbo bungo
//
// EXPECTED OUTPUT
//  There are 3 people!
//  Hello great bilbo !
//  Hello great balbo !
//  Hello great bungo !
//  Nice to meet you all.
// ---------------------------------------------------------

func ClArgs4a() {
	l := len(os.Args) - 1
	first := os.Args[1]
	second := os.Args[2]
	third := os.Args[3]
	fmt.Printf("There are %d people!\n", l)
	fmt.Printf("Hello great %s !\n", first)
	fmt.Printf("Hello great %s !\n", second)
	fmt.Printf("Hello great %s !\n", third)
	fmt.Println("Nice to meet you all.")

}

func ClArgs4b() {
	var (
		l  = len(os.Args) - 1
		n1 = os.Args[1]
		n2 = os.Args[2]
		n3 = os.Args[3]
	)

	fmt.Println("There are", l, "people!")
	fmt.Println("Hello great", n1, "!")
	fmt.Println("Hello great", n2, "!")
	fmt.Println("Hello great", n3, "!")
	fmt.Println("Nice to meet you all.")

}

// ---------------------------------------------------------
// EXERCISE: Greet 5 People
//
//  Greet 5 people this time.
//
//  Please do not copy paste from the previous exercise!
//
// RESTRICTION
//  This time do not use variables.
//
// INPUT
//  bilbo balbo bungo gandalf legolas
//
// EXPECTED OUTPUT
//  There are 5 people!
//  Hello great bilbo !
//  Hello great balbo !
//  Hello great bungo !
//  Hello great gandalf !
//  Hello great legolas !
//  Nice to meet you all.
// ---------------------------------------------------------

func ClArgs5() {
	fmt.Println("There are", len(os.Args)-1, "people!")
	fmt.Println("Hello great", os.Args[1])
	fmt.Println("Hello great", os.Args[2])
	fmt.Println("Hello great", os.Args[3])
	fmt.Println("Hello great", os.Args[4])
	fmt.Println("Hello great", os.Args[5])
	fmt.Println("Nice to meet you all.")
}

// ---------------------------------------------------------
// EXERCISE: Windows Path
//
//  1. Change the following program
//  2. It should use a raw string literal instead
//
// HINT
//  Run this program first to see its output.
//  Then you can easily understand what it does.
//
// EXPECTED OUTPUT
//  Your solution should output the same as this program.
//  Only that it should use a raw string literal instead.
// ---------------------------------------------------------

func MyString1() {
	// HINTS:
	// \\ equals to backslash character
	// \n equals to newline character
	/*
		path := "c:\\program files\\duper super\\fun.txt\n" +
				"c:\\program files\\really\\funny.png"
	*/

	path := `c:\program files\duper super\fun.txt
c:\program files\really\funny.png`

	fmt.Println(path)
}

// ---------------------------------------------------------
// EXERCISE: Print JSON
//
//  1. Change the following program
//  2. It should use a raw string literal instead
//
// HINT
//  Run this program first to see its output.
//  Then you can easily understand what it does.
//
// EXPECTED OUTPUT
//  Your solution should output the same as this program.
//  Only that it should use a raw string literal instead.
// ---------------------------------------------------------

func MyString2() {
	// HINTS:
	// \t equals to TAB character
	// \n equals to newline character
	// \" equals to double-quotes character

	json := "\n" +
		"{\n" +
		"\t\"Items\": [{\n" +
		"\t\t\"Item\": {\n" +
		"\t\t\t\"name\": \"Teddy Bear\"\n" +
		"\t\t}\n" +
		"\t}]\n" +
		"}\n"
	fmt.Println(json)

	json2 := `
{			
	"Items": [{
		"Item": {
			"name": "Teddy Bear"
		}
	}]
}`

	fmt.Println(json2)
}

// ---------------------------------------------------------
// EXERCISE: Raw Concat
//
//  1. Initialize the name variable
//     by getting input from the command line
//
//  2. Use concatenation operator to concatenate it
//     with the raw string literal below
//
// NOTE
//  You should concatenate the name variable in the correct
//  place.
//
// EXPECTED OUTPUT
//  Let's say that you run the program like this:
//   go run main.go inanç
//
//  Then it should output this:
//   hi inanç!
//   how are you?
// ---------------------------------------------------------

func MyStrings3() {
	name := os.Args[1]

	msg := `hi ` + name + ` !
how are you?`

	fmt.Println(msg)
}

// ---------------------------------------------------------
// EXERCISE: Count the Chars
//
//  1. Change the following program to work with unicode
//     characters.
//
// INPUT
//  "İNANÇ"
//
// EXPECTED OUTPUT
//  5
// ---------------------------------------------------------

func MyStrings4() {
	length := utf8.RuneCountInString(os.Args[1])
	fmt.Println(length)
}

// ---------------------------------------------------------
// EXERCISE: Improved Banger
//
//  Change the Banger program the work with Unicode
//  characters.
//
// INPUT
//  "İNANÇ"
//
// EXPECTED OUTPUT
//  İNANÇ!!!!!
// ---------------------------------------------------------

func MyStrings5() {
	msg := os.Args[1]

	s := msg + strings.Repeat("!", utf8.RuneCountInString(msg))

	fmt.Println(s)
}

// ---------------------------------------------------------
// EXERCISE: ToLowercase
//
//  1. Look at the documentation of strings package
//  2. Find a function that changes the letters into lowercase
//  3. Get a value from the command-line
//  4. Print the given value in lowercase letters
//
// HINT
//  Check out the strings package from Go online documentation.
//  You will find the lowercase function there.
//
// INPUT
//  "SHEPARD"
//
// EXPECTED OUTPUT
//  shepard
// ---------------------------------------------------------

func MyStrings6() {
	s := strings.ToLower(os.Args[1])
	fmt.Println(s)
}

// ---------------------------------------------------------
// EXERCISE: Trim It
//
//  1. Look at the documentation of strings package
//  2. Find a function that trims the spaces from
//     the given string
//  3. Trim the text variable and print it
//
// EXPECTED OUTPUT
//  The weather looks good.
//  I should go and play.
// ---------------------------------------------------------

func MyStrings7() {
	msg := `
	
	The weather looks good.
I should go and play.




	`

	fmt.Println(strings.TrimSpace(msg))
}

// ---------------------------------------------------------
// EXERCISE: Right Trim It
//
//  1. Look at the documentation of strings package
//  2. Find a function that trims the spaces from
//     only the right-most part of the given string
//  3. Trim it from the right part only
//  4. Print the number of characters it contains.
//
// RESTRICTION
//  Your program should work with unicode string values.
//
// EXPECTED OUTPUT
//  5
// ---------------------------------------------------------

func MyStrings8() {
	// currently it prints 19
	// it should print 7

	name := "  inanç           "
	fmt.Println(len(name))
	fmt.Println(utf8.RuneCountInString(strings.TrimRight(name, " ")))
}
