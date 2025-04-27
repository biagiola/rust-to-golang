package main

import "fmt"

func even_or_odd(numbers int) {
	result := ""

	if numbers%2 == 0 {
		result = "even"
	} else {
		result = "odd"
	}
	fmt.Println("The number is %s\n", result)
}

func main() {
	even_or_odd(17)
}
