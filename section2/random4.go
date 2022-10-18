package section2

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Verbose Mode
//
//  When the player runs the game like this:
//
//    go run main.go -v 4
//
//  Display each generated random number:

//    1 3 4 🎉  YOU WIN!
//
//  In this example, computer picks 1, 3, and 4. And the
//  player wins.
//
// HINT
//  You need to get and interpret the command-line arguments.
// ---------------------------------------------------------

const (
	usage4 = `Welcome to the Lucky Number Game! 🍀

The program will pick %d random numbers.

The greater your number is, the harder it gets.

Wanna play?`

	maxTurn4 = 5
)

func VerboseMode() {
	arg := os.Args[1:]

	if len(arg) < 1 {
		fmt.Printf(usage4, maxTurn4)
		return
	}

	guess, err := strconv.Atoi(arg[len(arg)-1]) //1 2 3 -> arg: {1, 2, 3} -> len(arg): 3 -> arg[2] = guess
	if err != nil {
		fmt.Println("Not a number")
		return
	}

	if guess <= 0 { // can not be 0 becaue of always matching
		fmt.Println("Give me a positive number")
		return
	}
	// fmt.Printf("Your number: %d.\n", guess)

	rand.Seed(time.Now().UnixNano()) // random number source

	var verbose bool
	if arg[0] == "-v" {
		verbose = true
	}

	for turn := 1; turn <= maxTurn4; turn++ {
		rn := rand.Intn(guess + 1)
		// fmt.Printf("#%d: Lucky number = %d\n", turn, rn)

		if verbose {
			fmt.Printf("%d ", rn)
		}

		if rn == guess {
			if turn == 1 {
				fmt.Println("🥇 FIRST TIME WINNER!!!")
			} else {
				fmt.Println("🎉  YOU WON!")
			}
			return
		}
	}
	/*
		// Better solution: Avoid nested statements.

		for turn := 1; turn <= maxTurns; turn++ {
			n := rand.Intn(guess) + 1

			// Better, why?
			//
			// Instead of nesting the if statement into
			//   another if statement; it simply continues.
			//
			// TLDR: Avoid nested statements.
			if n != guess {
				continue
			}

			if turn == 1 {
				fmt.Println("🥇 FIRST TIME WINNER!!!")
			} else {
				fmt.Println("🎉  YOU WON!")
			}
			return
		}
	*/

	fmt.Println("☠️ YOU LOST... Try again?")

}
