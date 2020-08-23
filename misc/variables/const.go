package main

import "fmt"

const x string = "Changes Never Come..."

/*
func constTest() {
	const x string = "Changes Never Come..."
	fmt.Println(x)
	x = "Changes?"
	fmt.Println(x)
}
*/

func main() {
	fmt.Println(x)
	x :=  "Changes?"
	fmt.Println(x)
}
