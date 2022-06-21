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
		fmt.Printf("%s.\n", os)

	}

	fmt.Println("When's Monday?")
	today := time.Now().Weekday()
	switch int(time.Monday) {
	case int(today) + 0:
		fmt.Println("Today.")
	case int(today) + 1:
		fmt.Println("Tomorrow.")
	case int(today) + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")

	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}

}
