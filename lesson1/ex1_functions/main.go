package main

import "fmt"

func main() {

	fmt.Println(square(3))
	fmt.Println(foo())
	fmt.Println(getNaked())
}

func square(i int) int {
	return i * i
}

func foo() (string, int) {
	return "foo", 42
}

func getNaked() (what string) {
	what = "foo"
	return
}
