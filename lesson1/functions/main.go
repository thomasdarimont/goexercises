package main

import "fmt"

func main() {

	fmt.Println(add(1, 2))

	a, b := swap("foo", "bar")
	fmt.Println(a, b)
}
func swap(a string, b string) (string, string) {
	return b, a
}

func add(a, b int) int {
	return a + b
}
