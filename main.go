package main

func main() {

	/*
		// Learn the basic of os.Args
		fmt.Printf("%#v\n", os.Args)

		fmt.Println("Path: ", os.Args[0])
		fmt.Println("1st argument: ", os.Args[1])
		fmt.Println("2nd argument: ", os.Args[2])
		fmt.Println("3rd argument: ", os.Args[3])

		fmt.Println("Number of items indside os.Args: ", len(os.Args))

	*/

	// EXERCISES

	/*
		exercises.Variables1()      // passed
		exercises.Variables2()      // passed
		exercises.Variables3()      // passed
		exercises.Variables4a()     // passed
		exercises.Variables4b(5, 6) // * passed
		exercises.Variables5()      // not pass: using a multiple assigment statement
		exercises.Variables6()      // * passed
		exercises.Variables7()      // * not pass: using the resul of the function to assign to the variables
		exercises.Variables8()      // passed
		exercises.Variables9()      // passed
		exercises.Variables10()     // passed
	*/

	/*
		exercises.TypeC1() // passed
		exercises.TypeC2() // passed
		exercises.TypeC3() // not passed: float64(5.5) -> 5.5
		exercises.TypeC4() // passed
		exercises.TypeC5() // passed
	*/

	// fmt.Printf("%#v\n", os.Args)
	// exercises.ClArgs1() // passed
	// exercises.ClArgs2() // passed
	// exercises.ClArgs3() //* passed
	// exercises.ClArgs4b() //** passed: how to fix when input <3 variables?
	// exercises.ClArgs5() // passed

	// fmt.Println(`C:\loctruong\github.com`)

	// name := "春天"
	// fmt.Println(len(name))                    // -> 6: len() actually returns the number of bytes
	// fmt.Println(utf8.RuneCountInString(name)) // returns acttuall the number of characters

	// myFunc()

	// exercises.MyString1() // passed
	// exercises.MyString2() // passed
	// exercises.MyStrings3() // not passed: `Hi ` + name + ` !`
	// exercises.MyStrings4() // passed
	// exercises.MyStrings5() // passed
	// exercises.MyStrings6() // passed
	// exercises.MyStrings7() // passed
	// exercises.MyStrings8() // * not passed: TrimRight(s, " ") not Trim

}

// func myFunc() {
// 	// s := os.Args[1]
// 	// e := strings.Repeat("!", len(s))
// 	// fmt.Println(strings.ToUpper(s) + e)

// 	s := os.Args[1]
// 	e := strings.Repeat("!", utf8.RuneCountInString(s))
// 	// fmt.Println(strings.ToUpper(s) + e)

// 	fmt.Println(e + strings.ToUpper(s) + e)

// }
