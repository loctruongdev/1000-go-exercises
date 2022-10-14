package main

import (
	exercisessection2 "1000go/exercisesection2"
)

const (
	maxTurns = 5 // less is more difficult
	usage    = `Welcome to the Lucky Number Game! 🍀
The program will pick %d random numbers.
Your mission is to guess one of those numbers.
The greater your number is, harder it gets.
Wanna play?
`
)

func main() {

	// exercisessection2.FirstTurnWinner()
	// exercisessection2.RandomMessages()
	// exercisessection2.DoubleGuesses()
	// exercisessection2.VerboseMode()
	// exercisessection2.EnoughPicks()
	exercisessection2.DynamicDiff()

}
