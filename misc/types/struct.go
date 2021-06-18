package main

import "fmt"

// Working with structs (similar to classes in Python)

type Person struct {
	name, hobby string
	age         int
}

func main() {
	newGuy := Person{name: "Joe", hobby: "skateboarding", age: 58}
	fmt.Println(newGuy)
	fmt.Println(newGuy.name)
	fmt.Println(newGuy.hobby)
	fmt.Println(newGuy.age)
	fmt.Printf("Hello, my name is %s. I love %s. I am %v years old.", newGuy.name, newGuy.hobby, newGuy.age)
}
