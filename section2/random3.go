package section2

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Double Guesses
//
//  Let the player guess two numbers instead of one.
//
// HINT:
//  Generate random numbers using the greatest number
//  among the guessed numbers.
//
// EXAMPLES
//  go run main.go 5 6
//  Player wins if the random number is either 5 or 6.
// ---------------------------------------------------------

const (
	maxTurn3 = 5 // less is more difficult
	usage3   = `Welcome to the Lucky Number Game! 🍀
The program will pick %d random numbers.
Your mission is to guess one of those numbers.
The greater your number is, harder it gets.
Wanna play?
`
)

func DoubleGuesses() {
	args := os.Args[1:]

	if len(args) != 2 && len(args) != 1 {
		fmt.Printf(usage3, maxTurn3)
	} else if len(args) == 1 {
		fmt.Println("We need two numbers")
		return
	}

	guess, err := strconv.Atoi(args[0])
	guess2, err2 := strconv.Atoi(args[1])

	if err != nil || err2 != nil {
		fmt.Println("Not a number.")
		return
	}

	if guess <= 0 || guess2 <= 0 {
		fmt.Println("Please pick positive numbers.")
		return
	}

	max := guess
	if guess < guess2 {
		max = guess2
	}

	rand.Seed(time.Now().UnixNano())

	for turn := 0; turn < maxTurn3; turn++ {
		n := rand.Intn(max + 1)
		fmt.Printf("Lucky number: %d \n", n)

		if n == guess || n == guess2 {
			fmt.Println("🎉  YOU WIN!")
			return
		}
	}

	fmt.Println("☠️  YOU LOST... Try again?")
}
