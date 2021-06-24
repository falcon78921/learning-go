package main

import "fmt"

// Create a fixed array of 10 elements that contain integers
var storage [10]int

// buildArray() will iterate over a for loop and add values to storage[]
func buildArray() [10]int {
	for i := 0; i <= 9; i++ {
		storage[i] = i
	}
	return storage
}

func main() {
	p := buildArray()
	fmt.Println(p)
	fmt.Println(p[0])
	fmt.Println(p[5])
}
