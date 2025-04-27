package main

import "fmt"

func main() {
	counter := 10

	for {
		if counter < 0 {
			fmt.Println("Blastoff!!")
			break
		}
		fmt.Printf("%d near to blastoff..!!\n", counter)
		counter--
	}
}
