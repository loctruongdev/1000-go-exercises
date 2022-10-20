package section3

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func RetroClock() {
	type placeholder [5]string

	zero := placeholder{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}

	one := placeholder{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}

	two := placeholder{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}

	three := placeholder{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}

	four := placeholder{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}

	five := placeholder{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}

	six := placeholder{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}

	seven := placeholder{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}

	eight := placeholder{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}

	nine := placeholder{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}

	digits := [...]placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	for line := range digits[0] {
		for digit := range digits {
			fmt.Printf(digits[digit][line] + " ")
		}
		fmt.Println()
	}

	separator := placeholder{
		"   ",
		" ░ ",
		"   ",
		" ░ ",
		"   ",
	}

	screen.Clear()

	for {

		h, m, s := time.Now().Hour(), time.Now().Minute(), time.Now().Second()
		// fmt.Println(h, m, s)

		clock := [...]placeholder{
			digits[h/10], digits[h%10],
			separator,
			digits[m/10], digits[m%10],
			separator,
			digits[s/10], digits[s%10],
		}

		screen.MoveTopLeft()

		for line := range clock[0] {
			for i, digit := range clock {
				data := clock[i][line]
				if digit == separator && s%2 == 0 {
					data = "   "
				}
				fmt.Printf(data + " ")

			}
			fmt.Println()
		}

		time.Sleep(time.Second)

	}

}
