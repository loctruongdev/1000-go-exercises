package theory

import (
	"1000go/theory"
)

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

	/*
		const (
			EST = -(5 + iota)
			_
			MST
			PST
		)

		fmt.Println(EST, MST, PST)
	*/

	// exercises.MyIOTA1() //passed
	// exercises.MyIOTA2() //passed
	// exercises.MyIOTA3() //* passed

	/*
		var brand string = "Google"
		fmt.Printf("%q\n", brand)
		fmt.Printf("%s\n", brand)
		fmt.Printf("%T\n", brand)
		fmt.Printf("%v\n", brand)
		fmt.Printf("%#v", brand)
	*/

	// Examples: PART I
	/*
		// I'm using multiple declarations instead of singular
		var (
			speed int
			heat  float64
			off   bool
			brand string
		)

		fmt.Printf("%T\n", speed) // %T -> type
		fmt.Printf("%T\n", heat)
		fmt.Printf("%T\n", off)
		fmt.Printf("%T\n", brand)
	*/

	// Examples: PART II

	/*
		var (
			planet   = "venus"
			distance = 261
			orbital  = 224.701
			hasLife  = false
		)

		// swiss army knife %v verb
		// fmt.Printf("Planet: %v\n", planet)
		// fmt.Printf("Distance: %v millions kms\n", distance)
		// fmt.Printf("Orbital Period: %v days\n", orbital)
		fmt.Printf("Does %v have life? %v\n", planet, hasLife)

		// argument indexing - indexes start from 1
		fmt.Printf(
			"%v is %v away. Think! %[2]v kms! %[1]v OMG.\n",
			planet, distance,
		)

		// why use other verbs than? because: type-safety
		// uncomment to see the warnings:
		//
		// fmt.Printf("Planet: %d\n", planet)
		// fmt.Printf("Distance: %s millions kms\n", distance)
		// fmt.Printf("Orbital Period: %t days\n", orbital)
		// fmt.Printf("Does %v has life? %f\n", planet, hasLife)
		fmt.Printf("Planet: %s\n", planet)                    // ok
		fmt.Printf("Distance: %d millions kms\n", distance)   //ok
		fmt.Printf("Orbital Period: %v days\n", orbital)      // %f -> float
		fmt.Printf("Does %v has life? %v\n", planet, hasLife) // %s | %t = true/false

		// correct verbs:
		fmt.Printf("Planet: %s\n", planet)
		fmt.Printf("Distance: %d millions kms\n", distance)
		fmt.Printf("Orbital Period: %f days\n", orbital)
		fmt.Printf("Does %s has life? %t\n", planet, hasLife)

		fmt.Println("====================")
		// precision
		fmt.Printf("Orbital Period: %f days\n", orbital) // orbital  = 224.701
		fmt.Printf("Orbital Period: %.0f days\n", orbital)
		fmt.Printf("Orbital Period: %.1f days\n", orbital)
		fmt.Printf("Orbital Period: %.2f days\n", orbital)
	*/

	// exercises.MyPrintf1() // passed
	// exercises.MyPrintf1b()
	// exercises.MyPrintf2()
	// exercises.MyPrintf2b() //* not pass:
	// exercises.MyPrintf3() // passed
	// exercises.MyPrintf4() // passed
	// exercises.MyPrintf5() // passed
	// exercises.MyPrintf6() // * not passed: use %[1]T
	// exercises.MyPrintf7() // passed
	// exercises.MyPrintf8() // passed
	// exercises.MyPrintf9() // passed
	// exercises.MyPrintf10() // passed

	/*
		var on bool
		on = !on
		fmt.Println(on)

		on = !!on
		fmt.Println(on)
	*/

	// exercises.MyIf1()
	// exercises.MyIf2a()
	// exercises.MyIf2b()
	// exercises.MyIf3()
	// exercises.MyIf3b()
	// exercises.MyIf3c()
	// exercises.MyIf4()
	// exercises.MyIf5()
	// exercises.MyF5b()
	// exercises.MyIf6()

	// exercises.FeettoMeters()
	// exercises.MyErrHandling1()
	// exercises.MyErrHandling2()
	// exercises.LeapYear()
	// exercises.LeapYear2()
	// exercises.LeapYear3()
	// exercises.DaysofMonth()

	// exercises.RichterScales()
	// exercises.RichterScales2()
	// exercises.RichterScales2b()
	// exercises.UseSwitch()
	// exercises.MyConvert()
	// exercises.MyConvert2()

	/*
		city := os.Args[1]

		switch city {
		case "paris", "paris2":
			{
				fmt.Println("France")
			}
		case "tokyo":
			{
				fmt.Println("Japan")
			}
		default:
			fmt.Println("Vietnam")
		}
	*/

	/*

		if len(os.Args) != 2 {
			fmt.Println("Give me a month name.")
			return
		}

		switch m := os.Args[1]; m {
		case "Dec", "Jan", "Feb":
			{
				fmt.Println("Winter")
			}
		case "Mar", "Apr", "May":
			{
				fmt.Println("Spring")
			}
		case "Jun", "Jul", "Aug":
			{
				fmt.Println("Summer")
			}
		case "Sep", "Oct", "Nov":
			{
				fmt.Println("Fall")
			}
		default:
			fmt.Printf("%q is not a month.\n", m)
		}

	*/

	// fmt.Println(strings.IndexAny("aeiuo", os.Args[1]))
	// fmt.Println(strings.IndexAny("aeiuo", "a"))
	// fmt.Println(os.Args[1])
	// // fmt.Println(strings.IndexAny("crwth", "aeiouy"))
	// var sum int
	// for i := 1; i <= 1000; i++ {
	// 	sum += i
	// }
	// fmt.Println(sum)

	theory.MyLoop1()

}

// func myFunc() {
// 	// s := os.Args[1]
// 	// e := strings.Repeat("!", len(s))
// 	// fmt.Println(strings.ToUpper(s) + e)

// 	s := os.Args[1]
// 	e := strings.Repeat("!", utf8.RuneCountInString(s))
// 	// fmt.Println(strings.ToUpper(s) + e)

// 	fmt.Println(e + strin;gs.ToUpper(s) + e)
