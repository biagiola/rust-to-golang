package main

import "fmt"

func factorial_recursive(current int) int {
	if current <= 1 {
		return 1
	} else {
		return current * factorial_recursive(current-1)
	}
}

func main() {
	result := factorial_recursive(4)
	fmt.Printf("Factorial of four: %d", result)
}

// factorial (4) -> 4 * 3 * 2 * 1
// factorial (4) -> 4 * factorial(3)
// factorial (3) -> 3 * factorial(2)