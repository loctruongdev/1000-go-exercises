package main

import (
	"1000go/section4"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	/*
		var (
			namesA []string
		)

		var namesB = []string{}

		namesC := []string{}

		namesD := make([]string, 0)

		namesE := make([]string, 1)

		fmt.Printf("namesA: %T %v %d %t\n", namesA, namesA, len(namesA), namesA == nil)
		fmt.Printf("namesB: %T %v %d %t\n", namesB, namesB, len(namesB), namesB == nil)
		fmt.Printf("namesB: %T %v %d %t\n", namesC, namesC, len(namesC), namesC == nil)
		fmt.Printf("namesB: %T %v %d %t\n", namesD, namesD, len(namesD), namesD == nil)
		fmt.Printf("namesB: %T %v %d %t\n", namesE, namesE, len(namesE), namesE == nil)

	*/

	/*
		sliceA2 := append(sliceA, 4, 5, 6, 7)
		fmt.Println("sliceA2: ", sliceA2)

		// change value element 1
		sliceA2[0] = 10

		fmt.Println("sliceA2: ", sliceA2)

		// Checking sliceA
		fmt.Println("sliceA: ", sliceA) // wont effect
	*/

	// sliceA := make([]int, 0, 5)
	// sliceA = append(sliceA, 1, 2, 3)
	// fmt.Println("sliceA: ", sliceA)

	// sliceB := append(sliceA, 4)
	// fmt.Println("sliceB: ", sliceB)

	// sliceB[0] = 10

	// fmt.Println("sliceB: ", sliceB)
	// fmt.Println("sliceA: ", sliceA)

	// sliceB := append(sliceA, 3)
	// fmt.Println(sliceB)

	// sliceA := make([]int, 8)
	// for i := 0; i < 8; i++ {
	// 	sliceA[i] = i
	// }
	// fmt.Println(sliceA)
	// fmt.Printf("length: %d cap %d\n", len(sliceA), cap(sliceA))

	// sliceB := sliceA[0:4]
	// fmt.Printf("length: %d cap %d\n", len(sliceB), cap(sliceB))

	// for i := 0; i < 8; i++ {
	// 	sliceB[i] = i
	// }

	// ages := []int{10, 20, 30}
	// fmt.Println(ages)
	// ages = ages[0:3]
	// fmt.Println(ages)
	// ages = ages[0:2]
	// fmt.Println(ages)

	// sliceA := []string{"one", "two", "three", "four", "five"}
	// fmt.Println(sliceA)

	// sliceA1 := sliceA[0:3]
	// fmt.Println(sliceA1)

	// sliceA1[0] = "1"
	// fmt.Println(sliceA1)
	// fmt.Println(sliceA)

	// sliceA := []string{"one", "two", "three"}
	// fmt.Println(sliceA)

	// sliceB := []string{"4", "5", "6"}
	// fmt.Println(sliceB)

	// sliceC := append(sliceA, sliceB...)
	// fmt.Println(sliceC)
	// sliceC[0] = "NEW"
	// fmt.Println(sliceC)
	// fmt.Println(sliceA)

	// sliceD := sliceC[2:4]
	// fmt.Println(sliceD)

	// sliceD[0] = "THREE"
	// sliceD[1] = "FOUR"
	// fmt.Println(sliceD)
	// fmt.Println(sliceC)
	// fmt.Println(sliceA)
	// fmt.Println(sliceB)

	// sliceB := append(sliceA, "four")
	// fmt.Println(sliceA)
	// fmt.Println(sliceB)

	// SliceAcopy := sliceA
	// SliceAcopy = append(SliceAcopy, "five", "six")
	// fmt.Println(sliceA)
	// fmt.Println(SliceAcopy)

	// sliceC := append(sliceA, sliceB...)
	// fmt.Println(sliceC)
	// sliceCnew := sliceC[2:5]
	// fmt.Println(sliceCnew)

	/*

		screen.Clear()

		for {
			screen.MoveTopLeft()

			now := time.Now()
			h, m, s := now.Hour(), now.Minute(), now.Second()
			ss := now.Nanosecond() / 1e8

			clock := [...]r.Placeholder{
				r.Digits[h/10], r.Digits[h%10],
				r.Separator,
				r.Digits[m/10], r.Digits[m%10],
				r.Separator,
				r.Digits[s/10], r.Digits[s%10],
				r.Dot,
				r.Digits[ss],
			}

			// for line := range clock[0] {
			// 	for i, digit := range clock {
			// 		data := clock[i][line]
			// 		if digit == r.Separator && s%2 == 0 {
			// 			data = "   "
			// 		}
			// 		fmt.Printf(data + " ")

			// 	}
			// 	fmt.Println()
			// }

			// alarmed := s%10 == 0

			for line := range clock[0] {
				// if alarmed {
				// 	clock = r.Alarm
				// }

				for index, digit := range clock {
					next := clock[index][line]
					if (digit == r.Separator || digit == r.Dot) && s%2 == 0 {
						next = "   "
					}
					fmt.Print(next, "  ")
				}
				fmt.Println()
			}

			time.Sleep(time.Second / 10)

		}

	*/

	/*
		moods := [...][3]string{
			{"good", "awesome", "happy"},
			{"bad", "terrible", "worry"},
		}
		// fmt.Println(moods)

		args := os.Args[1:]
		if len(args) != 2 {
			fmt.Println("[your name] [positive | negative]")
			return
		}

		name, mood := args[0], args[1]

		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(len(moods[0]))

		// switch mood {
		// case "positive":
		// 	fmt.Printf("%s feels %s\n", name, moods[0][n])
		// case "negative":
		// 	fmt.Printf("%s feels %s\n", name, moods[1][n])
		// }

		// Alternative way
		var mi int
		if mood != "positive" {
			mi = 1
		}

		fmt.Printf("%s feels %s\n", name, moods[mi][n])
	*/

	/*
		moods := [...]string{
			"sad",
			"happy",
			"terrible",
			"awesome",
		}
		// fmt.Println(feelings)

		arg := os.Args[1:]
		fmt.Println(len(arg))

		if len(arg) != 1 {
			fmt.Println("Your Name")
			return
		}

		name := arg[0]

		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(len(moods))
		fmt.Printf("%s feels %s\n", name, moods[n])
	*/

	// ★★★★★ SECTION 1 ★★★★★

	// === VARIABLES EXERCISES ===

	// section1.MakeItBlue()
	// section1.VariablesToVariables()
	// section1.AssignWithExpressions()
	// section1.RectanglePerimeter()
	// section1.RectanglePerimeter2()
	// section1.MultiAssign()
	// section1.MultiAssignb()
	// section1.MultiShortFunction()
	// section1.MySwapper()
	// section1.MySwapper2()
	// section1.DiscardFiles()

	// === TYPE-CONVERSION EXERCISES ===

	// section1.TypeConversion()
	// section1.TypeConversion2()
	// section1.TypeConversion3()
	// section1.TypeConversion4()
	// section1.TypeConversion5()

	// === COMMAND LINE ARGUMENT EXERCISES ===

	// section1.Args1()
	// section1.Args2()
	// section1.Args3()
	// section1.Args4()
	// section1.Args4b()
	// section1.Args5()

	// === STRINGS EXERCISES ===

	// section1.WindowsPath()
	// section1.PrintJSON()
	// section1.RawConcat()
	// section1.CountChars()
	// section1.ImprovedBanger()
	// section1.ToLowerCase()
	// section1.TrimIt()
	// section1.RightTrimIt()

	// === IOTA EXERCISES ===

	// section1.IotaMonths()
	// section1.IotaMonths2()
	// section1.IotaSeasons()

	// === PRINTF EXERCISES ===

	// section1.PrintYourAge()
	// section1.PrintNameLastname()
	// section1.FalseClaims()
	// section1.PrintTemperature()
	// section1.DoubleQuotes()
	// section1.PrintType()
	// section1.PrintType2()
	// section1.PrintType3()
	// section1.PrintType4()
	// section1.PrintFullname()

	// === IF STATEMENT EXERCISES ===

	// section1.AgeSeasons()
	// section1.SimplifyIt()
	// section1.SimplifyItb()
	// section1.ArgCount()
	// section1.ArgCountb()
	// section1.ArgCountc()
	// section1.VowelOrConsonant()
	// section1.VowelOrConsonantb()
	// section1.VowelOrConsonantc()
	// section1.Challenge1()
	// section1.Challenge1b()
	// section1.Challenge2()

	// === ERR HANDLING EXERCISES ===

	// section1.MovieRatings()
	// section1.OddOrEven()
	// section1.LeapYear()
	// section1.LeapYearb()
	// section1.SimplifyLeapYear()
	// section1.DaysInMonth()
	// section1.DaysInMonthb()

	// === SWITCH EXERCISES ===

	// section1.RichterScales()
	// section1.RichterScales2()
	// section1.RichterScales2b()
	// section1.ConvertIf()
	// section1.ConvertSwitch()
	// section1.StringManipulator()
	// section1.StringManipulatorb()
	// section1.DaysInMonth2()

	// === LOOP EXERCISES ===

	// section1.DynamicTable()
	// section1.MathTables()
	// section1.SumNumbers()
	// section1.SumNumbersVerbose()
	// section1.SumN()
	// section1.SumNb()
	// section1.OnlyEvens()
	// section1.BreakUp()
	// section1.InfiniteKill()

	// ★★★★★ SECTION 2 ★★★★★

	// === RANDOMIZATION EXERCISES ===

	// section2.FirstTurnWinner()
	// section2.RandomMessages()
	// section2.DoubleGuesses()
	// section2.VerboseMode()
	// section2.EnoughPicks()
	// section2.DynamicDifficulty()

	// === LABELED STATEMENT EXERCISES ===

	// section2.CaseInsensitiveSearch()
	// section2.PathSearcher()
	// section2.PathSearcherb()
	// section2.CrunchPrimes()
	// section2.CrunchPrimesb()

	// ★★★★★ SECTION 3 ★★★★★

	// === ARRAYS EXERCISES ===
	// section3.DeclareEmptyArrays()
	// section3.GetSetArrayElements()
	// section3.GetSetArrayElements()
	// section3.RefactorArray()
	// section3.RefactorEllipsis()
	// section3.FixCode()
	// section3.CompareArrays()
	// section3.AssignArrays()
	// section3.WizardPrinter()
	// section3.WizardPrinterb()
	// section3.CurrencyConverter()
	// section3.BookstoreSearchEngine()
	// section3.BookstoreSearchEngineb()
	// section3.FindAverage()
	// section3.FindAverageb()
	// section3.NumberSorter()
	// section3.NumberSorter()
	// section3.NumberSorterb()
	// section3.WordFinder()
	// section3.RetroClock()
	// retroclock.Clock()

	// ★★★★★ SECTION 4 ★★★★★

	// === SLICES EXERCISES ===
	// section4.DeclareSlices()
	// section4.DeclareSlicesb()
	// section4.DeclareSlices2()
	// section4.AssignSliceLiterals()
	// section4.ArrayToSlice()
	// section4.FixProblems()
	// section4.CompareSlices()
	// section4.Append()
	// section4.Append2()
	// section4.Append3()
	// section4.AppendAndSort()
	// section4.HousingPrices()
	// section4.HousingPricesAverages()
	// section4.HousingPricesAveragesb()
	// section4.SliceNumbers()
	// section4.SliceArgs()
	// section4.SliceArgsb()
	// section4.SliceHousePrices()
	// section4.PriceHousePricesDA()
	// section4.FixBackingArray()
	// section4.FixBackingArrayDA()
	// section4.SortBackingArray()
	// section4.ObserveMemoryAllocations()
	// section4.ObserveLengthCapacity()
	// section4.ObserveCapacityGrowth()
	// section4.CorrectLyric()
	// section4.CorrectLyricDA()
	// section4.PracticeAdvanceSlice()
	// section4.LimitBackingArraySharing()
	// section4.FixMemoryLeak()
	// section4.AddNewLineEachSentenceDA()
	// section4.PrintDailyRequest()
	section4.PrintDailyRequestDA()

	// words := []string{"lucy", "in", "the", "sky", "with", "diamonds"}
	// fmt.Printf("words1 %v\n", words)
	// words = append(words[:1], "is", "everywhere")
	// fmt.Printf("words2 %v\n", words)
	// words = words[:6]
	// fmt.Printf("words3 %v\n", words)
	// words = append(words, words[len(words)+1:cap(words)]...)
	// fmt.Printf("words3 %v\n", words)

	// words := []string{1022: ""}
	// fmt.Printf("len: %d cap: %d\n", len(words), cap(words))
	// words = append(words, "boom!")
	// fmt.Printf("len2: %d cap2: %d\n", len(words), cap(words))

	// slice1 := make([]string, 3)
	// s.Show("slice1 before append: ", slice1)
	// slice1 = append(slice1, "one", "two")
	// slice1 = slice1[:cap(slice1)]
	// s.Show("slice1 after append: ", slice1)

	// slice2 := make([]string, 0, 3)
	// s.Show("slice2 before append: ", slice2)
	// slice2 = append(slice2, "one", "two")
	// slice2 = slice2[:cap(slice2)]
	// s.Show("slice2 affter append: ", slice2)

}

func init() {
	s.PrintBacking = true // prints the backing array
	s.MaxPerLine = 10     // prints 10 slice elements per line
	s.Width = 60          // prints 60 character per line
}
