package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	ifStatementDemo()
	switchStatementDemo()
	switchStatementWithNoCondition()

	deferStatementDemo()
}
func deferStatementDemo() {
	defer fmt.Println("hates fish")
	fmt.Println("Charlie")
}

func switchStatementWithNoCondition() {

	// end
	// Without a condition
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func switchStatementDemo() {
	// see also https://github.com/golang/go/wiki/Switch
	fmt.Print("Go runs on...")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OSX.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s\n", os)
	}
}

func ifStatementDemo() {
	var inbasket bool
	oranges := 1
	if inbasket {
		// Do stuff
	}
	if oranges != 1 {
		// Do stuff
	}
	// Short hand
	if i := 1; i < oranges {
		// Do stuff
		// We can only use i here though
	}
	// Cant use i here
}
