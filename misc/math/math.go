package main

import "fmt"

func math1() {
	fmt.Println("1 + 1 =", 1 + 1)
}

func math2() {
	fmt.Println("1 + 1 =", 1.0 + 1.0)
}

func main() {
	math1()
	math2()
}
