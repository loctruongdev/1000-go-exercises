package theorysection1b

import (
	"1000go/exercises"
	"fmt"
	"os"
	"strconv"
)

func main() {

	/*
		var (
			sum int
			i   = 1
		)

		for {
			if i > 10 {
				break
			}

			if i%2 != 0 {
				i++
				continue
			}

			sum += i
			fmt.Println(i, "-->", sum)
			i++
		}

		fmt.Println(sum)
	*/

	/*
		const max = 5

		fmt.Printf("%5s", "X")
		for i := 0; i <= max; i++ {
			fmt.Printf("%5d", i)
		}
		fmt.Println()

		for i := 0; i <= max; i++ {
			fmt.Printf("%5d", i)

			for j := 0; j <= max; j++ {
				fmt.Printf("%5d", i*j)
			}

			fmt.Println()
		}
	*/

	// MyTableMulti()
	// MyTablePlus()

	// MyTableDiv()
	// MyTableMinus()

	// exercises.MyLoopEx1()
	// exercises.MyLoopEx2()
	// exercises.MySum()
	// exercises.MySum2()
	// exercises.MySum3()
	// exercises.MySum3b()
	// exercises.MySum4()
	// exercises.MySum5()
	exercises.MySum6()

	/*
		if len(os.Args) != 3 {
			fmt.Println("gimme two numbers")
			return
		}

		min, err1 := strconv.Atoi(os.Args[1])
		max, err2 := strconv.Atoi(os.Args[2])
		if err1 != nil || err2 != nil || min >= max {
			fmt.Println("wrong numbers")
			return
		}

		var (
			i   = min
			sum int
		)

		for {
			if i > max {
				break
			} else if i%2 != 0 {
				i++
				continue
			}

			fmt.Print(i)
			if i < max-1 {
				fmt.Print(" + ")
			}

			sum += i
			i++
		}
		fmt.Printf(" = %d\n", sum)
	*/

}

func MyTableMulti() {

	if len(os.Args) != 2 {
		fmt.Println("Give me a valid number")
		return
	}

	max, err := strconv.Atoi(os.Args[1])

	if err != nil || max <= 0 {
		fmt.Println("Give me a valid number")
		return
	}

	fmt.Printf("%5s", "X")
	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)

		for j := 0; j <= max; j++ {
			fmt.Printf("%5d", i*j)
		}

		fmt.Println()
	}

}

func MyTablePlus() {
	if len(os.Args) != 2 {
		fmt.Println("Give me a valid number")
		return
	}

	max, err := strconv.Atoi(os.Args[1])

	if err != nil || max <= 0 {
		fmt.Println("Give me a valid number")
		return
	}

	fmt.Printf("%5s", "X")
	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)

		for j := 0; j <= max; j++ {
			fmt.Printf("%5d", i+j)
		}

		fmt.Println()
	}
}

func MyTableDiv() {
	if len(os.Args) != 2 {
		fmt.Println("Give me a valid number")
		return
	}

	max, err := strconv.Atoi(os.Args[1])

	if err != nil || max <= 0 {
		fmt.Println("Give me a valid number")
		return
	}

	fmt.Printf("%5s", "X")
	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)

		fmt.Printf("%5s", "-")

		for j := 1; j <= max; j++ {
			fmt.Printf("%5d", i/j)
		}

		fmt.Println()
	}
}

func MyTableMinus() {
	if len(os.Args) != 2 {
		fmt.Println("Give me a valid number")
		return
	}

	max, err := strconv.Atoi(os.Args[1])

	if err != nil || max <= 0 {
		fmt.Println("Give me a valid number")
		return
	}

	fmt.Printf("%5s", "X")
	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)

		for j := 0; j <= max; j++ {
			fmt.Printf("%5d", i-j)
		}

		fmt.Println()
	}

}
