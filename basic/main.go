package main

import "fmt"

// func multiply(a int, b int) int {
// 	return a + b
// }

// func whoAmI(name string, age int) string {
// 	// this is like string template literals but we do need to pass a placeholder in each values we want to be in a variable.
// 	return fmt.Sprintf("Hello %s, you are %d old", name, age)
// }

func main() {
	// var age int = multiply(20, 3) // explicit (manual declaration)
	// name:= "rahman" // inferred (the great walrus operator xd)

	// fmt.Println("Hello My Name is:", name, "My age is:", age) // normal log
	// msg := whoAmI("rahman", 23)
	// fmt.Println(msg)

	const (
		test = 1 << iota // 0001
		execute // 0010
		final // 0100
		anotherVal // 1000
	)

	fmt.Println(test, execute, final)
}