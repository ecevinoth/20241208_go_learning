package main

import "fmt"

func main() {
	// variable declaration
	var name string = "John"
	var age int = 20
	var isStudent bool = true
	var height float64 = 175.5
	var pi float32 = 3.14159
	fmt.Println(name, age, isStudent, height, pi)

	// polluated declaration
	// if you declare a variable by not useing var keyword, it is called polluated declaration
	a, b, c, _, d := 1, 2, 3, 4, 5 // _ is blank identifier not used variable
	fmt.Println(a, c, d)           // ERROR:         00_helloworld\01_variables.go:16:5: other declaration of b

	// zero value
	var z int     // 0
	var s string  // ""
	var b bool    // false
	var f float32 // 0.0
	fmt.Println(z, s, b, f)

	// short hand declaration
	firstName := "John"
	lastName := "Doe"
	dateOfBirth := 19901231

	fmt.Println(firstName, lastName, dateOfBirth)

}
