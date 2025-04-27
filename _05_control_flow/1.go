package main

import "fmt"

func main() {
	number := 8

	switch number {
	case 2, 4, 6, 8:
		fmt.Printf("%d is even\n", number)
	case 1, 3, 5, 7:
		fmt.Printf("%d is odd\n", number)
	default:
		fmt.Println("Unknow number")
	}

	// short version
	switch {
	case number%2 == 0:
		fmt.Printf("%d is even\n", number)
	default:
		fmt.Printf("%d is odd\n", number)
	}
}

// we need to use commas instead of or symbol -> |
// in go, the or symbol is doing a "bitwise OR" operation
// where we make a kind of comparartion of the two numbers
// in its binary form, following this rules:
// . if either bit is 1, the result is 1
// . if both bits are 0, result is 0
// So... a code in go like this:
// case 2 | 4 | 6 | 8: // == case 14
// will thrown us 14 in our case options, getting in the
// default arm.
// why 14?
// because we doing the bitwise or one by be take us to this:
// 2 | 4 -> 0011 | 0100 = 0110 // this is six in binary, move to the next bitwise or
// 6 | 6 -> 0110 | 0110 = 0110 // give us six again
// 6 | 9 -> 0110 | 1000 = 1110 // and this is 14.

// So, in rust, the or symbol "|"" is just a comparation, a logical math operation,
// but in go is the 'bitwise or' operation.
