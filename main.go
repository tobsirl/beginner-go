package main

import (
	"fmt"
)

// Structs
type Person struct {
	Name    string
	Age     int
	Country string
}

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

	// Control structures
	// if-else
	if age < 18 {
		fmt.Println("You are a minor.")
	} else if age >= 18 && age < 65 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a senior citizen.")
	}

	day := "Wednesday"

	// switch
	switch day {
	case "Monday":
		fmt.Println("Today is Monday.")
	case "Tuesday":
		fmt.Println("Today is Tuesday.")
	case "Wednesday":
		fmt.Println("Today is Wednesday.")
	case "Thursday":
		fmt.Println("Today is Thursday.")
	case "Friday":
		fmt.Println("Today is Friday.")
	case "Saturday":
		fmt.Println("Today is Saturday.")
	case "Sunday":
		fmt.Println("Today is Sunday.")
	default:
		fmt.Println("Invalid day.")
	}

	// for loop
	for i := 1; i <= 5; i++ {
		fmt.Printf("Iteration %d\n", i)
	}

	// while loop (using for)
	i := 1
	for i <= 5 {
		fmt.Printf("While loop iteration %d\n", i)
		i++
	}

	// Arrays
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array of numbers:", numbers)

	// Slices
	allNumbers := numbers[:]   // Slicing the array
	firstThree := numbers[0:3] // Slicing the first three elements
	fmt.Println("First three numbers:", firstThree)
	fmt.Println("Slice of numbers:", allNumbers)
	// Appending to a slice
	allNumbers = append(allNumbers, 6, 7, 8)
	fmt.Println("Appended slice of numbers:", allNumbers)

	fruits := []string{}
	fruits = append(fruits, "Apple", "Banana", "Cherry")
	fmt.Println("Slice of fruits:", fruits)
	moreFruits := []string{"Mango", "Orange"}
	fruits = append(fruits, moreFruits...)
	fmt.Println("Extended slice of fruits:", fruits)

	// Maps
	capitalCities := map[string]string{
		"USA":    "Washington, D.C.",
		"Canada": "Ottawa",
		"Mexico": "Mexico City",
	}
	// Accessing map values
	fmt.Println("Capital of USA:", capitalCities["USA"])
	// Adding a new key-value pair
	capitalCities["Germany"] = "Berlin"
	fmt.Println("Capital cities:", capitalCities)
	capitalCities["France"] = "Paris"
	// Iterating over a map
	for country, capital := range capitalCities {
		fmt.Printf("The capital of %s is %s.\n", country, capital)
	}
	// Deleting a key-value pair
	delete(capitalCities, "Mexico")
	fmt.Println("Updated capital cities:", capitalCities)

	// Structs
	person := Person{
		Name:    "Alice",
		Age:     28,
		Country: "Canada",
	}
	fmt.Printf("Person: Name: %s, Age: %d, Country: %s\n", person.Name, person.Age, person.Country)
	// Accessing struct fields
	fmt.Println("Person's name:", person.Name)

}

func add(a int, b int) int {
	return a + b
}

func calculateSumAndProduct(a, b int) (int, int) {
	sum := a + b
	product := a * b
	return sum, product
}
