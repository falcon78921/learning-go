package main

import "fmt"

const x string = "Changes Never Come..."

func constTest() {
	x := "Changes?"
	fmt.Println(x)
}

func main() {
        constTest()
	fmt.Println(x)
	x := "Changes?"
	fmt.Println(x)
}
