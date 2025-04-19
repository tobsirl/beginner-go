package main

import "fmt"

func main() {
	// Variable declaration
	var name string = "John Doe"
	fmt.Printf("Hello, %s!\n", name)

	age := 30
	fmt.Printf("You are %d years old.\n", age)

	var city string
	city = "New York"
	fmt.Printf("You live in %s.\n", city)

	var country, continent string = "USA", "North America"
	fmt.Printf("You live in %s, %s.\n", country, continent)

	var (
		profession string = "Software Engineer"
		salary     int    = 100000
		isEmployed bool   = true
	)
	fmt.Printf("You are a %s with a salary of $%d. Employed: %t\n", profession, salary, isEmployed)

	// Zero values
	var defaultInt int
	var defaultFloat float64
	var defaultString string
	var defaultBool bool
	fmt.Printf("Zero values: int: %d, float: %f, string: '%s', bool: %t\n", defaultInt, defaultFloat, defaultString, defaultBool)

	// Constants
	const pi float64 = 3.14
	fmt.Printf("Value of pi: %f\n", pi)

	const (
		Monday    = iota // 0
		Tuesday          // 1
		Wednesday        // 2
		Thursday         // 3
		Friday           // 4
		Saturday         // 5
		Sunday           // 6
	)
	fmt.Printf("Days of the week: %d, %d, %d, %d, %d, %d, %d\n", Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday)

	// functions
	result := add(5, 10)
	fmt.Printf("The sum of 5 and 10 is: %d\n", result)

	sum, product := calculateSumAndProduct(3, 4)
	fmt.Printf("The sum of 3 and 4 is: %d, and the product is: %d\n", sum, product)
}

func add(a int, b int) int {
	return a + b
}

func calculateSumAndProduct(a, b int) (int, int) {
	sum := a + b
	product := a * b
	return sum, product
}
