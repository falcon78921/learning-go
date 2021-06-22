// Build a person (practicing concepts in golang)

package main

import "fmt"

// Create a Person data structure
type Person struct {
	name, hobby, status string
	age                 int
}

// createPerson() takes an input and returns a Person struct
func createPerson(name string) *Person {

	var status string
	var hobby string
	var age int

	// Look at the value for name and process accordingly
	switch name {
	case "Joey":
		status = "Single"
		hobby = "Playing baseball"
		age = 58
	case "John":
		status = "Married"
		hobby = "Rocking out"
		age = 35
	case "Jimmy":
		status = "Single"
		hobby = "Playing guitar"
		age = 25
	case "Devon":
		status = "Married"
		hobby = "Reading comics"
		age = 27
	default:
		fmt.Println("Enter a valid name.")
	}

	// Invoke the Person struct and fill out according to name
	invokePerson := Person{name: name, status: status, hobby: hobby, age: age}
	// Return a Person struct when finished
	return &invokePerson
}

func main() {
	p := createPerson("Joey")
	fmt.Println(p)
	fmt.Println(p.age)
}
