package section7

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
	"github.com/mattn/go-runewidth"
)

// ---------------------------------------------------------
// EXERCISE: Adjust the width and height automatically
//
//    Instead of setting the width and height manually,
//    you need to get the width and height of the terminal
//    screen from your operating system.
//
//  1. Update your program to use my screen package.
//     It offers an easy way to get the width and height.
//
//     go get -u https://github.com/inancgumus/screen
//
//  2. Read the package's documentation and find a way to
//     get the screen size: width and height.
//
//     The documentation is here:
//     https://godoc.org/github.com/inancgumus/screen
//
//  3. Use it to set the board's dimensions.
//
//
// OPTIONAL EXERCISE
//
//  1. When you set the width, you may see that the ball
//     goes beyond the left and right borders. This happens
//     because the ball emoji spans to multiple console
//     columns (or cells). Ordinary characters have a
//     single column.
//
//     1. Get the width of the ball emoji using a function
//        from the following package:
//
//        go get -u github.com/mattn/go-runewidth
//
//     2. Divide the width using the rune width of the
//        ball emoji.
//
//  2. Your terminal may have borders, so reduce the
//     height by taking into account the height of
//     your terminal borders.
//
//
// EXPECTED OUTPUT
//
//  When you run the program, the ball should start
//  animating across the total width and height of your
//  terminal screen dynamically.
//
//  Currently you set width and height manually, so it
//  wasn't matter whether your terminal was bigger or
//  smaller, but now it will be!
//
// ---------------------------------------------------------

func BouncingBallEx2() {
	const (
		// width  = 50
		// height = 10

		cellEmpty = ' '
		cellBall  = '⚾'

		maxFrames = 1200
		speed     = time.Second / 20

		// drawing buffer length
		//
		// *2 for extra spaces
		// +1 for newlines

	)

	var (
		// width, height int
		px, py int    // ball position
		vx, vy = 1, 1 // velocities

		cell rune // current cell (for caching)
	)

	//get dimension
	width, height := screen.Size()
	fmt.Printf("Width: %d Height: %d\n", width, height)

	ballWidth := runewidth.RuneWidth(cellBall)

	width /= ballWidth
	height--
	fmt.Printf("Width2: %d Height2: %d\n", width, height)

	// create the board
	board := make([][]bool, width)
	for column := range board {
		board[column] = make([]bool, height)
	}

	// create a drawing buffer
	bufLen := (width*2 + 1) * height
	buf := make([]rune, 0, bufLen)

	// clear the screen once
	screen.Clear()

	for i := 0; i < maxFrames; i++ {
		// calculate the next ball position
		px += vx
		py += vy

		// when the ball hits a border reverse its direction
		if px <= 0 || px >= width-1 {
			vx *= -1
		}
		if py <= 0 || py >= height-1 {
			vy *= -1
		}

		// remove the previous ball
		for y := range board[0] {
			for x := range board {
				board[x][y] = false
			}
		}

		// put the new ball
		board[px][py] = true

		// rewind the buffer (allow appending from the beginning)
		buf = buf[:0]

		// draw the board into the buffer
		for y := range board[0] {
			for x := range board {
				cell = cellEmpty
				if board[x][y] {
					cell = cellBall
				}
				buf = append(buf, cell, ' ')
			}
			buf = append(buf, '\n')
		}

		// print the buffer
		screen.MoveTopLeft()
		fmt.Print(string(buf))

		// slow down the animation
		time.Sleep(speed)
	}
}

func BouncingBallEx2DA() {
	const (
		cellEmpty = ' '
		cellBall  = '⚾'

		maxFrames = 1200
		speed     = time.Second / 20
	)

	var (
		px, py int    // ball position
		vx, vy = 1, 1 // velocities

		cell rune // current cell (for caching)
	)

	// you can get the width and height using the screen package easily:
	width, height := screen.Size()

	// get the rune width of the ball emoji
	ballWidth := runewidth.RuneWidth(cellBall)

	// adjust the width and height
	width /= ballWidth
	height-- // there is a 1 pixel border in my terminal

	// create the board
	board := make([][]bool, width)
	for column := range board {
		board[column] = make([]bool, height)
	}

	// drawing buffer length
	// *2 for extra spaces
	// +1 for newlines
	bufLen := (width*2 + 1) * height

	// create a drawing buffer
	buf := make([]rune, 0, bufLen)

	// clear the screen once
	screen.Clear()

	for i := 0; i < maxFrames; i++ {
		// calculate the next ball position
		px += vx
		py += vy

		// when the ball hits a border reverse its direction
		if px <= 0 || px >= width-1 {
			vx *= -1
		}
		if py <= 0 || py >= height-1 {
			vy *= -1
		}

		// remove the previous ball
		for y := range board[0] {
			for x := range board {
				board[x][y] = false
			}
		}

		// put the new ball
		board[px][py] = true

		// rewind the buffer (allow appending from the beginning)
		buf = buf[:0]

		// draw the board into the buffer
		for y := range board[0] {
			for x := range board {
				cell = cellEmpty
				if board[x][y] {
					cell = cellBall
				}
				buf = append(buf, cell, ' ')
			}
			buf = append(buf, '\n')
		}

		// print the buffer
		screen.MoveTopLeft()
		fmt.Print(string(buf))

		// slow down the animation
		time.Sleep(speed)
	}
}
