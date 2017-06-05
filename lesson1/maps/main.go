package main

import "fmt"

type position struct {
	lat, lng float64
}

func main() {
	names := make(map[string]string)
	names["charlie"] = "smith"
	fmt.Println(names)

	// You can also use your own types in maps
	people := make(map[string]position)
	people["charlie"] = position{
		lat: 53.0000,
		lng: 54.4000,
	}
	fmt.Println(people)

	// You can check whether a key exists in a map with two value assignment
	v, ok := people["charlie"]
	// v will return the value if exists and ok will be a boolean
	fmt.Println("The value:", v, "Present?", ok)

	// You can check whether a key exists in a map with two value assignment
	vb, ok := people["bobby"]
	if ok {
		fmt.Println("Present", "The value:", vb, "Present?", ok)
	} else {
		fmt.Println("Not present", "The value:", vb, "Present?", ok)
	}

}
