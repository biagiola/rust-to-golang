package main

import "fmt"

func countdown(seconds int) {
	if seconds == 0 {
		fmt.Println("Blastoff..!!")
	} else {
		fmt.Printf("%d seconds to blastoff..!!\n", seconds)
		countdown(seconds - 1)
	}
}

func main() {
	countdown(5)
}
