package main

import "fmt"

func main() {
	season := "summer"

	if season == "summer" {
		fmt.Println("School's out")
	} else if season == "winter" {
		fmt.Println("Brr, so cold")
	} else {
		fmt.Println("Lots of rain!")
	}
}

// any difference with rust this case