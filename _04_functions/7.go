package main

import "fmt"

func main() {
	season := "winter"

	switch season {
	case "summer":
		fmt.Println("School's out")
	case "winter":
		fmt.Println("Brr, so cold")
	default:
		fmt.Println("Lot's of rain")
	}
}
