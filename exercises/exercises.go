package exercises

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
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

// ---------------------------------------------------------
// EXERCISE: Iota Months
//
//  1. Initialize the constants using iota.
//  2. You should find the correct formula for iota.
//
// RESTRICTIONS
//  1. Remove the initializer values from all constants.
//  2. Then use iota once for initializing one of the
//     constants.
//
// EXPECTED OUTPUT
//  9 10 11
// ---------------------------------------------------------

func MyIOTA1() {
	const (
		Nov = 11 - iota
		Oct
		Sep
	)

	fmt.Println(Sep, Oct, Nov)
}

// ---------------------------------------------------------
// EXERCISE: Iota Months #2
//
//  1. Initialize multiple constants using iota.
//  2. Please follow the instructions inside the code.
//
// EXPECTED OUTPUT
//  1 2 3
// ---------------------------------------------------------

func MyIOTA2() {
	// HINT: This is a valid constant declaration
	//       Blank-Identifier can be used in place of a name
	// const _ = iota
	//    ^- this is just a name

	// Now, use iota and initialize the following constants
	// "automatically" to 1, 2, and 3 respectively.
	const (
		_ = iota
		Jan
		Feb
		Mar
	)

	// This should print: 1 2 3
	// Not: 0 1 2
	fmt.Println(Jan, Feb, Mar)
}

// ---------------------------------------------------------
// EXERCISE: Iota Seasons
//
//  Use iota to initialize the season constants.
//
// HINT
//  You can change the order of the constants.
//
// EXPECTED OUTPUT
//  12 3 6 9
// ---------------------------------------------------------

func MyIOTA3() {
	// NOTE : You should remove all the initializers below
	//        first. Then use iota to fix it.
	const (
		Spring = 3 * (iota + 1) // 3*(0+1)
		Summer                  //3*(1+1)
		Fall                    //3*(2+1)
		Winter                  //3*(3+1)
	)

	fmt.Println(Winter, Spring, Summer, Fall)
}

// ---------------------------------------------------------
// EXERCISE: Print Your Age
//
//  Print your age using Printf
//
// EXPECTED OUTPUT
//  I'm 30 years old.
//
// NOTE
//  You should change 30 to your age, of course.
// ---------------------------------------------------------

func MyPrintf1() {
	age := 30
	fmt.Printf("I'm %d years old\n", age)
}

func MyPrintf1b() {
	fmt.Printf("I'm %d years old\n", 35)
}

// ---------------------------------------------------------
// EXERCISE: Print Your Name and LastName
//
//  Print your name and lastname using Printf
//
// EXPECTED OUTPUT
//  My name is Inanc and my lastname is Gumus.
//
// BONUS
//  Store the formatting specifier (first argument of Printf)
//    in a variable.
//  Then pass it to printf
// ---------------------------------------------------------

func MyPrintf2() {
	firstName, lastName := "Spring", "Truong"
	fmt.Printf("My name is %s and my last name is %s\n.", firstName, lastName)
}

func MyPrintf2b() {
	// fmt.Printf("My name is %s and my last name is %s\n.", "Spring", "Truong")

	// BONUS
	msg := "My name is %s and my last name is %s.\n"
	fmt.Printf(msg, "Spring", "Truong")
}

// ---------------------------------------------------------
// EXERCISE: False Claims
//
//  Use printf to print the expected output using a variable.
//
// EXPECTED OUTPUT
//  These are false claims.
// ---------------------------------------------------------

func MyPrintf3() {
	tf := false
	fmt.Printf("These are %t claims.\n", tf)
}

// ---------------------------------------------------------
// EXERCISE: Print the Temperature
//
//  Print the current temperature in your area using Printf
//
// NOTE
//  Do not use %v verb
//  Output "shouldn't" be like 29.500000
//
// EXPECTED OUTPUT
//  Temperature is 29.5 degrees.
// ---------------------------------------------------------

func MyPrintf4() {
	temp := 29.5
	fmt.Printf("Temperature is %.1f degrees.\n", temp)
}

// ---------------------------------------------------------
// EXERCISE: Double Quotes
//
//  Print "hello world" with double-quotes using Printf
//  (As you see here)
//
// NOTE
//  Output "shouldn't" be just: hello world
//
// EXPECTED OUTPUT
//  "hello world"
// ---------------------------------------------------------

func MyPrintf5() {
	s := "hello world"
	fmt.Printf("%q\n", s)
}

// ---------------------------------------------------------
// EXERCISE: Print the Type
//
//  Print the type and value of 3 using fmt.Printf
//
// EXPECTED OUTPUT
//  Type of 3 is int
// ---------------------------------------------------------

func MyPrintf6() {
	fmt.Printf("Type of %d is %[1]T\n", 3)
}

// ---------------------------------------------------------
// EXERCISE: Print the Type #2
//
//  Print the type and value of 3.14 using fmt.Printf
//
// EXPECTED OUTPUT
//  Type of 3.14 is float64
// ---------------------------------------------------------

func MyPrintf7() {
	fmt.Printf("Type of %.2f is %[1]T\n", 3.14)
}

// ---------------------------------------------------------
// EXERCISE: Print the Type #3
//
//  Print the type and value of "hello" using fmt.Printf
//
// EXPECTED OUTPUT
// 	Type of hello is string
// ---------------------------------------------------------

func MyPrintf8() {
	fmt.Printf("Type of %s is %[1]T\n", "hello")
}

// ---------------------------------------------------------
// EXERCISE: Print the Type #4
//  Print the type and value of true using fmt.Printf
//
// EXPECTED OUTPUT
//  Type of true is bool
// ---------------------------------------------------------

func MyPrintf9() {
	fmt.Printf("Type of %t is %[1]T\n", true)
}

// ---------------------------------------------------------
// EXERCISE: Print Your Fullname
//
//  1. Get your name and lastname from the command-line
//  2. Print them using Printf
//
// EXAMPLE INPUT
//  Inanc Gumus
//
// EXPECTED OUTPUT
//  Your name is Inanc and your lastname is Gumus.
// ---------------------------------------------------------

func MyPrintf10() {
	fmt.Printf("Your name is %s and your last name is %s.\n", os.Args[1], os.Args[2])
}

// ---------------------------------------------------------
// EXERCISE: Age Seasons
//
//  Let's start simple. Print the expected outputs,
//  depending on the age variable.
//
// EXPECTED OUTPUT
//  If age is greater than 60, print:
//    Getting older
//  If age is greater than 30, print:
//    Getting wiser
//  If age is greater than 20, print:
//    Adulthood
//  If age is greater than 10, print:
//    Young blood
//  Otherwise, print:
//    Booting up
// ---------------------------------------------------------

func MyIf1() {
	age := 10
	if age > 60 {
		fmt.Println("Getting older")
	} else if age > 30 {
		fmt.Println("Getting wiser")
	} else if age > 20 {
		fmt.Println("Adulthood")
	} else if age > 10 {
		fmt.Println("Young blood")
	} else {
		fmt.Println("Booting up")
	}
}

// ---------------------------------------------------------
// EXERCISE: Simplify It
//
//  Can you simplify the if statement inside the code below?
//
//  When:
//    isSphere == true and
//    radius is equal or greater than 200
//
//    It will print "It's a big sphere."
//
//    Otherwise, it will print "I don't know."
//
// EXPECTED OUTPUT
//  It's a big sphere.
// ---------------------------------------------------------

func MyIf2a() {
	// DO NOT TOUCH THIS
	isSphere, radius := true, 200

	var big bool

	if radius >= 50 {
		if radius >= 100 {
			if radius >= 200 {
				big = true
			}
		}
	}

	if big != true {
		fmt.Println("I don't know.")
	} else if !(isSphere == false) {
		fmt.Println("It's a big sphere.")
	} else {
		fmt.Println("I don't know.")
	}
}

func MyIf2b() {
	// DO NOT TOUCH THIS
	isSphere, radius := true, 200

	if isSphere && radius >= 200 {
		fmt.Println("It's a big sphere.")
	} else {
		fmt.Println("I don't know.")
	}

}

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

func MyIf3() {
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

func MyIf3b() {

	args := []string{"minh", "loc"}
	var msg string
	for i := 0; i <= len(args)-1; i++ {
		msg = msg + " " + args[i]
	}
	fmt.Printf("There is argument(s): %s\n", msg)

}

func MyIf3c() {
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

// ---------------------------------------------------------
// EXERCISE: Vowel or Consonant
//
//  Detect whether a letter is vowel or consonant.
//
// NOTE
//  y or w is called a semi-vowel.
//  Check out: https://www.merriam-webster.com/words-at-play/why-y-is-sometimes-a-vowel-usage
//  Check out: https://www.dictionary.com/e/w-vowel/
//
// HINT
//  + You can find the length of an argument using the len function.
//
//  + len(os.Args[1]) will give you the length of the 1st argument.
//
// BONUS
//  Use strings.IndexAny function to detect the vowels.
//  Search on Google for: golang pkg strings IndexAny
//
// Furthermore, you can also use strings.ContainsAny. Check out: https://golang.org/pkg/strings/#ContainsAny
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me a letter
//
//  go run main.go hey
//    Give me a letter
//
//  go run main.go a
//    "a" is a vowel.
//
//  go run main.go y
//    "y" is sometimes a vowel, sometimes not.
//
//  go run main.go w
//    "w" is sometimes a vowel, sometimes not.
//
//  go run main.go x
//    "x" is a consonant.
// ---------------------------------------------------------

// len(os.Args) == 1 || len(os.Args) >2 -> Give me a letter
// len(os.Args) == 2 -> IndexAny("aeiuo", os.Args[1]) -> os.Args[1] is a vowel |
// IndexAny("yw", os.Args[1]) -> os.Args[1] is sometimes a vowel, sometimes not.
// else: os.Args[1] is a consonant.

func MyIf4() {
	if len(os.Args) == 2 && len(os.Args[1]) == 1 {
		c := strings.IndexAny("aeiuo", os.Args[1])
		if c != -1 {
			fmt.Printf("%q is a vowel.\n", os.Args[1])
		} else {
			c := strings.IndexAny("yw", os.Args[1])
			if c != -1 {
				fmt.Printf("%q is sometimes a vowel, sometimes not.\n", os.Args[1])
			} else {
				fmt.Printf("%q is a consonant.\n", os.Args[1])
			}
		}
	} else {
		fmt.Println("Give me a letter.")
	}
}

func MyIf4b() {
	args := os.Args

	if len(args) != 2 || len(args[1]) != 1 {
		fmt.Println("Give me a letter")
		return
	}

	// I didn't use a short-if here because, it's already
	// hard to read. Do not make it harder.

	s := args[1]
	if s == "a" || s == "e" || s == "i" || s == "o" || s == "u" {
		fmt.Printf("%q is a vowel.\n", s)
	} else if s == "w" || s == "y" {
		fmt.Printf("%q is sometimes a vowel, sometimes not.\n", s)
	} else {
		fmt.Printf("%q is a consonant.\n", s)
	}
}

func MyIf4c() {
	args := os.Args

	if len(args) != 2 || len(args[1]) != 1 {
		fmt.Println("Give me a letter")
		return
	}

	s := args[1]
	if strings.IndexAny(s, "aeiou") != -1 {
		fmt.Printf("%q is a vowel.\n", s)
	} else if s == "w" || s == "y" {
		fmt.Printf("%q is sometimes a vowel, sometimes not.\n", s)
	} else {
		fmt.Printf("%q is a consonant.\n", s)
	}

	// Notice that:
	//
	// I didn't use IndexAny for the else if above.
	//
	// It's because, calling a function is a costly operation.
	// And, this way, the code is simpler.
}

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

func MyIf5() {
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

func MyF5b() {
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

// ---------------------------------------------------------
// CHALLENGE #2
//  Add one more user to the PassMe program below.
//
// EXAMPLE USERS
//  username: jack
//  password: 1888
//
//  username: inanc
//  password: 1879
//
// EXPECTED OUTPUT
//  go run main.go
//    Usage: [username] [password]
//
//  go run main.go hacker 42
//    Access denied for "hacker".
//
//  go run main.go jack 1888
//    Access granted to "jack".
//
//  go run main.go inanc 1879
//    Access granted to "inanc".
//
//  go run main.go jack 1879
//    Invalid password for "jack".
//
//  go run main.go inanc 1888
//    Invalid password for "inanc".
// ---------------------------------------------------------

const (
	usage       = "Usage: [username] [password]"
	errUser     = "Access denied for %q.\n"
	errPwd      = "Invalid password for %q.\n"
	accessOK    = "Access granted to %q.\n"
	user, user2 = "jack", "inanc"
	pass, pass2 = "1888", "1879"
)

func MyIf6() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println(usage)
		return
	}

	u, p := os.Args[1], os.Args[2]

	if u != user && u != user2 {
		fmt.Printf(errUser, u)
	} else if u == user && p == pass {
		fmt.Printf(accessOK, u)
	} else if u == user2 && p == pass2 {
		fmt.Printf(accessOK, u)
	} else {
		fmt.Printf(errPwd, p)
	}
}

//FEET TO METERs

func FeettoMeters() {
	arg := os.Args[1]
	feet, err := strconv.ParseFloat(arg, 64)

	if err != nil {
		fmt.Printf("error: %q is not a number.\n", arg)
		return
	}

	meter := feet * 0.3048
	fmt.Printf("%g feet is %g meters.\n", feet, meter)

}

// ---------------------------------------------------------
// STORY
//
//  Your boss wants you to create a program that will check
//  whether a person can watch a particular movie or not.
//
//  Imagine that another program provides the age to your
//  program. Depending on what you return, the other program
//  will issue the tickets to the person automatically.
//
// EXERCISE: Movie Ratings
//
//  1. Get the age from the command-line.
//
//  2. Return one of the following messages if the age is:
//     -> Above 17         : "R-Rated"
//     -> Between 13 and 17: "PG-13"
//     -> Below 13         : "PG-Rated"
//
// RESTRICTIONS:
//  1. If age data is wrong or absent let the user know.
//  2. Do not accept negative age.
//
// BONUS:
//  1. BONUS: Use if statements only twice throughout your program.
//  2. BONUS: Use an if statement only once.
//
// EXPECTED OUTPUT
//  go run main.go 18
//    R-Rated
//
//  go run main.go 17
//    PG-13
//
//  go run main.go 12
//    PG-Rated
//
//  go run main.go
//    Requires age
//
//  go run main.go -5
//    Wrong age: "-5"
// ---------------------------------------------------------

// age from the common-line -> os.Args
// Requires age
// Wrong age: "-5"
// -> Above 17         : "R-Rated"
//     -> Between 13 and 17: "PG-13"
//     -> Below 13         : "PG-Rated"

func MyErrHandling1() {
	length := len(os.Args) - 1
	if length != 1 {
		fmt.Println("Requires age.")
		return
	}

	age, err := strconv.Atoi(os.Args[1])

	if err != nil || age < 0 {
		fmt.Printf("Wrong age: %q.\n", os.Args[1])
		return // -> exit error!
	} else if age > 17 {
		fmt.Println("R-Rated")
	} else if age <= 17 && age >= 13 {
		fmt.Println("PG-Rated")
	} else {
		fmt.Println("PG-13")
	}

}

// ---------------------------------------------------------
// EXERCISE: Odd or Even
//
//  1. Get a number from the command-line.
//
//  2. Find whether the number is odd, even and divisible by 8.
//
// RESTRICTION
//  Handle the error. If the number is not a valid number,
//  or it's not provided, let the user know.
//
// EXPECTED OUTPUT
//  go run main.go 16
//    16 is an even number and it's divisible by 8
//
//  go run main.go 4
//    4 is an even number
//
//  go run main.go 3
//    3 is an odd number
//
//  go run main.go
//    Pick a number
//
//  go run main.go ABC
//    "ABC" is not a number
// ---------------------------------------------------------

// Pick a number
// "ABC" is not a number
// Even | Even & div 8 | Odd

func MyErrHandling2() {
	if len(os.Args) != 2 {
		fmt.Println("Pick a number")
		return
	}

	n, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Printf("%q is not a number.\n", os.Args[1])
		return
	}

	if n%8 == 0 {
		fmt.Printf("%d is an even number and it's divisible by 8 \n", n)
	} else if n%2 == 0 {
		fmt.Printf("%d is an even number \n", n)
	} else {
		fmt.Printf("%d is an odd number \n", n)
	}

}

// ---------------------------------------------------------
// EXERCISE: Leap Year
//
//  Find out whether a given year is a leap year or not.
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me a year number
//
//  go run main.go eighties
//    "eighties" is not a valid year.
//
//  go run main.go 2018
//    2018 is not a leap year.
//
//  go run main.go 2100
//    2100 is not a leap year.
//
//  go run main.go 2019
//    2019 is not a leap year.
//
//  go run main.go 2020
//    2020 is a leap year.
//
//  go run main.go 2024
//    2024 is a leap year.
// ---------------------------------------------------------

func LeapYear() {

	if len(os.Args) != 2 {
		fmt.Println("Give me a year number")
		return
	}

	y, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Printf("%q is not a valid year.\n", os.Args[1])
		return
	}

	if y%4 == 0 {
		if y%100 == 0 {
			if y%400 == 0 {
				fmt.Printf("%d is a leap year.\n", y)
			} else {
				fmt.Printf("%d is NOT a leap year.\n", y)
			}
		} else {
			fmt.Printf("%d is a leap year.\n", y)
		}
	} else {
		fmt.Printf("%d is NOT a leap year.\n", y)
	}
}

func LeapYear2() {
	if len(os.Args) != 2 {
		fmt.Println("Give me a year number")
		return
	}

	year, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("%q is not a valid year.\n", os.Args[1])
		return
	}

	var leap bool
	if year%400 == 0 {
		leap = true
	} else if year%100 == 0 {
		leap = false
	} else if year%4 == 0 {
		leap = true
	}

	if leap {
		fmt.Printf("%d is a leap year.\n", year)
	} else {
		fmt.Printf("%d is not a leap year.\n", year)
	}
}

func LeapYear3() {
	if len(os.Args) != 2 {
		fmt.Println("Give me a year number")
		return
	}

	year, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("%q is not a valid year.\n", os.Args[1])
		return
	}

	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		fmt.Printf("%d is a leap year.\n", year)
	} else {
		fmt.Printf("%d is not a leap year.\n", year)
	}
}

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

func DaysofMonth() {

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

func DaysofMonth2() {
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

// ---------------------------------------------------------
// STORY
//  You're curious about the richter scales. When reporters
//  say: "There's been a 5.5 magnitude earthquake",
//
//  You want to know what that means.
//
//  So, you've decided to write a program to do that for you.
//
// EXERCISE: Richter Scale
//
//  1. Get the earthquake magnitude from the command-line.
//  2. Display its corresponding description.
//
// ---------------------------------------------------------
// MAGNITUDE                    DESCRIPTION
// ---------------------------------------------------------
// Less than 2.0                micro
// 2.0 and less than 3.0        very minor
// 3.0 and less than 4.0        minor
// 4.0 and less than 5.0        light
// 5.0 and less than 6.0        moderate
// 6.0 and less than 7.0        strong
// 7.0 and less than 8.0        major
// 8.0 and less than 10.0       great
// 10.0 or more                 massive
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me the magnitude of the earthquake.
//
//  go run main.go ABC
//    I couldn't get that, sorry.
//
//  go run main.go 0.5
//    0.50 is micro
//
//  go run main.go 2.5
//    2.50 is very minor
//
//  go run main.go 3
//    3.00 is minor
//
//  go run main.go 4.5
//    4.50 is light
//
//  go run main.go 5
//    5.00 is moderate
//
//  go run main.go 6
//    6.00 is strong
//
//  go run main.go 7
//    7.00 is major
//
//  go run main.go 8
//    8.00 is great
//
//  go run main.go 11
//    11.00 is massive
// ---------------------------------------------------------

func RichterScales() {
	if len(os.Args) != 2 {
		fmt.Println("Give me the magnitude of the earthquake.")
		return
	}

	richter, err := strconv.ParseFloat(os.Args[1], 64)

	if err != nil {
		fmt.Println("I couldn't get that, sorry.")
		return
	}

	var desc string

	switch r := richter; {
	case r < 2.0:
		{
			desc = "micro"
		}
	case r >= 2.0 && r < 3.0:
		{
			desc = "micro minor"
		}
	case r >= 3.0 && r < 4.0:
		{
			desc = "minor"
		}
	case r >= 4.0 && r < 5.0:
		{
			desc = "light"
		}
	case r >= 5.0 && r < 6.0:
		{
			desc = "moderate"
		}
	case r >= 6.0 && r < 7.0:
		{
			desc = "strong"
		}
	case r >= 5.0 && r < 6.0:
		{
			desc = "moderate"
		}
	case r >= 7.0 && r < 8.0:
		{
			desc = "major"
		}
	case r >= 8.0 && r < 10.0:
		{
			desc = "great"
		}
	case r >= 10:
		{
			desc = "massive"
		}
	default:
		desc = "Give me the magnitude of the earthquake."

	}

	fmt.Printf("%.2f is %s.\n", richter, desc)

}

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

func UseIf() {
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

func UseSwitch() {
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

// ---------------------------------------------------------
// STORY
//  You want to write a program that will manipulate a
//  given string to uppercase, lowercase, and title case.
//
// EXERCISE: String Manipulator
//
//  1. Get the operation as the first argument.
//
//  2. Get the string to be manipulated as the 2nd argument.
//
// HINT
//  You can find the manipulation functions in the strings
//  package Go documentation (ToLower, ToUpper, Title).
//
// EXPECTED OUTPUT
//
//  go run main.go
//    [command] [string]
//
//    Available commands: lower, upper and title
//
//  go run main.go lower 'OMG!'
//    omg!
//
//  go run main.go upper 'omg!'
//    OMG!
//
//  go run main.go title "mr. charles darwin"
//    Mr. Charles Darwin
//
//  go run main.go genius "mr. charles darwin"
//    Unknown command: "genius"
// ---------------------------------------------------------

func MyConvert() {
	if len(os.Args) != 3 {
		fmt.Println("[command] [string]")
		return
	}

	// arg := os.Args[1]
	// str := os.Args[2]
	cmd, str := os.Args[1], os.Args[2]

	switch cmd {
	case "lower":
		fmt.Println(strings.ToLower(str))
	case "upper":
		fmt.Println(strings.ToUpper(str))
	case "title":
		fmt.Println(strings.Title(str))
	default:
		fmt.Printf("Unknown command: %q\n", cmd)
	}

}

func MyConvert2() {
	if len(os.Args) != 2 {
		fmt.Println("Give me a month name")
		return
	}

	year := time.Now().Year()
	leap := year%4 == 0 && (year%100 != 0 || year%400 == 0)

	days, month := 28, os.Args[1]

	switch strings.ToLower(month) {
	case "april", "june", "september", "november":
		days = 30
	case "january", "march", "may", "july", "august", "october", "december":
		days = 31
	case "february":
		if leap {
			days = 29
		}
	default:
		fmt.Printf("%q is not a month.\n", month)
		return // -> will not print next fmt.Printf

	}

	fmt.Printf("%q has %d days.\n", month, days)

}
