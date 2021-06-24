package main

import "fmt"

const s string = "stable"

func main() {
	fmt.Println(s)
	const n = 5
	const d = 3.5 / n
	fmt.Println(d)
}
