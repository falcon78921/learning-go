package main

import (
	"fmt"
	"reflect"
)

var a, b, c string

func main() {
	a = "Michael"
	b = "Joe"
	c = "Steve"

	if a == "Mike" {
		fmt.Println(reflect.TypeOf(a))
		fmt.Println("It's Mike!")
	} else if a == "Joe" {
		fmt.Println("It's not Mike!")
	} else {
		fmt.Println("Nothing...")
		fmt.Println(reflect.TypeOf(a))
	}
}
