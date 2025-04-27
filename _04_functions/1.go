package main

import "fmt"

func main() {
	// here, we declare and infer its type
	result := square(5)
	fmt.Println(result)
}

func square(number int) int {
	return number * number
}

// we can declare a variable and put it its type also, but it's not neccesary
// var result int = square(5)
