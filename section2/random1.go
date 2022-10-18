package section2

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: First Turn Winner
//
//  If the player wins on the first turn, then display
//  a special bonus message.
//
// RESTRICTION
//  The first picked random number by the computer should
//  match the player's guess.
//
// EXAMPLE SCENARIO
//  1. The player guesses 6
//  2. The computer picks a random number and it happens
//     to be 6
//  3. Terminate the game and print the bonus message
// ---------------------------------------------------------

const (
	usage1 = `Welcome to the Lucky Number Game! 🍀

The program will pick %d random numbers.

The greater your number is, the harder it gets.

Wanna play?`

	maxTurn1 = 5
)

func FirstTurnWinner() {
	arg := os.Args[1:]

	if len(arg) != 1 {
		fmt.Printf(usage1, maxTurn1)
		return
	}

	guess, err := strconv.Atoi(arg[0])
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

	for turn := 1; turn <= maxTurn1; turn++ {
		rn := rand.Intn(guess + 1)
		// fmt.Printf("#%d: Lucky number = %d\n", turn, rn)

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
