package main

import "fmt"

func main() {
	const name, age = "Kim", 22
	fmt.Printf("%s is %d years old.\n", name, age)
	fmt.Printf("%v is %v years old.\t %T and %T", name, age, name, age)
	fmt.Printf(`Hello world 
	welcome to the world of Go
	🐱😾😿🙀🐈`)
}
