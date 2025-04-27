package main

import "fmt"

func main() {
	evaluation := false

	switch evaluation {
	case true:
		fmt.Println("The value is true")
	case false:
		fmt.Println("The value is false")
	}

	var result int
	switch evaluation {
	case true:
		result = 20
	case false:
		result = -40
	}
	fmt.Printf("Result: %d\n", result)
}

package main

import "fmt"

func main() {
	evaluation := false

	switch evaluation {
	case true:
		fmt.Println("The value is true")
	case false:
		fmt.Println("The value is false")
	}

	var result int
	switch evaluation {
	case true:
		result = 20
	case false:
		result = -40
	}
	fmt.Printf("Result: %d\n", result)
}

// In Rust, match is an expression (returns a value).
// In Go, switch is a statement (it does things, but doesn't return a value directly).

// That's why in Go we have to assign manually inside case blocks.
