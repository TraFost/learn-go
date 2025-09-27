package main

import "fmt"

func multiply(a int, b int) int {
	return a + b
}

func main() {
	var age int = multiply(20, 3) // explicit (manual declaration)
	name:= "rahman" // inferred (the great walrus operator xd)

	fmt.Println("Hello My Name is:", name, "My age is:", age)
}