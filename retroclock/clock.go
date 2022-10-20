package retroclock

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func Clock() {

	screen.Clear()

	for {
		screen.MoveTopLeft()

		now := time.Now()
		h, m, s := now.Hour(), now.Minute(), now.Second()
		ss := now.Nanosecond() / 1e8

		clock := [...]placeholder{
			digits[h/10], digits[h%10],
			separator,
			digits[m/10], digits[m%10],
			separator,
			digits[s/10], digits[s%10],
			dot,
			digits[ss],
		}

		// for line := range clock[0] {
		// 	for i, digit := range clock {
		// 		data := clock[i][line]
		// 		if digit == Separator && s%2 == 0 {
		// 			data = "   "
		// 		}
		// 		fmt.Printf(data + " ")

		// 	}
		// 	fmt.Println()
		// }

		// alarmed := s%10 == 0

		for line := range clock[0] {
			// if alarmed {
			// 	clock = Alarm
			// }

			for index, digit := range clock {
				next := clock[index][line]
				if (digit == separator || digit == dot) && s%2 == 0 {
					next = "   "
				}
				fmt.Print(next, "  ")
			}
			fmt.Println()
		}

		time.Sleep(time.Second / 10)

	}
}
