package main

import "fmt"

func main() {

	fmt.Println(add(1, 2))

	a, b := swap("foo", "bar")
	fmt.Println(a, b)

	x, y := split(22)
	fmt.Println(x, y)
}
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // naked return
}

func swap(a string, b string) (string, string) {
	return b, a
}

func add(a, b int) int {
	return a + b
}
