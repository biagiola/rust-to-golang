package main

import "fmt"

func main() {
	apply_to_jobs(4, "Sofware development")

	result := is_even(-2)
	fmt.Println(result)

	find_letter("polyndrome", 'y')
}

func apply_to_jobs(number int, title string) {
	fmt.Printf("I'm applying to %d %s jobs \n", number, title)
}

func is_even(number int) bool {
	return number%2 == 0
}

func find_letter(word string, letter rune) {
	for _, c := range word {
		if c == letter {
			fmt.Printf("%c", c)
		}
	}
}
