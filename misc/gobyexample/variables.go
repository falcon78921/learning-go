package main

import (
	"fmt"
	"reflect"
)

func main() {
	var floater float32
	a := "Hello, World!"
	b := 50
	floater = 1.05
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(floater))
}
