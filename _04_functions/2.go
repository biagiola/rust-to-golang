package main

import "fmt"

type Student struct {
	Name   string
	Age    int
	Active bool
}

func main() {
	// normal variables
	student_name := "David Biagiola"
	student_age := 99
	student_is_active := true
	fmt.Println(student_name, student_age, student_is_active)

	// to simulate rust tuples, we can use structs
	student := Student{"David Biagiola", 31, true}
	fmt.Println(student)

	// empty struct (equivalent to a unit, that is the most simple typle in rust)
	var empty struct{}
	fmt.Println(empty)

	// an empty function returns nothing
	nothingness := mystery()
	fmt.Println(nothingness)
}

func mystery() string {
	fmt.Println("Whataver")
	return "" // we need to return something if we want to create a variable
	// in the upper variable called unit.
}

// side notes: we cannot create unused variables
