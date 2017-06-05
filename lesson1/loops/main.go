package main

import "fmt"

func main() {

	//basic for loop
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	// go's while loop
	counter := 0
	for counter < 10 {
		counter++
		fmt.Println("counter", counter)
	}

	// for over range
	var array = map[string]int{"apples": 1, "oranges": 20, "bananas": 4}
	for key, value := range array {
		fmt.Println(key)
		fmt.Println(value)
	}

	// "infinite" loop

	counter = 0
	for {
		counter++
		if counter > 10 {
			break
		}

		fmt.Println("counter", counter)
	}
}
