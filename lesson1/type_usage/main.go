package main

import (
	"fmt"
	"strconv"
)

const doyoulovego = true

var (
	firstname string = "Charlie"
	surname   string = "Smith"
	age       int    = 22
)

func main() {
	if firstname == "Charlie" && doyoulovego == true {
		fmt.Println(firstname + `is ` + strconv.Itoa(age) + ` and loves Go`)
	}
}
