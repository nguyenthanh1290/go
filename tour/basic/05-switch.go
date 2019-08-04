package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Print("Go runs on ")

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
	fmt.Println("\n")

	// switch evaluation order
	today := time.Now().Weekday()
	fmt.Println("Today is: ", today)

	fmt.Println("When's Saturday?")

	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
	fmt.Println("\n")

	// Switch with no condition
	t := time.Now()
	fmt.Println("Now is: ", t)
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

// Go only runs the selected case, not all the cases that follow.
// In effect, the break statement that is needed at the end of each case in those languages is provided automatically in Go.

// Another important difference is that Go's switch cases need not be constants,
// and the values involved need not be integers.
