package main

import "fmt"

var names [3]string

func slicing() {
	slices := make([]int, 4)
	fmt.Println(slices)
	for i := 0; i <= 3; i++ {
		slices[i] = i
	}
	fmt.Println(slices)
	fmt.Println(len(slices))
	new := append(slices, 4, 5, 6)
	fmt.Println(new)
	fmt.Println(len(new))
	// Copy two slices
	q := make([]int, 11)
	for i := 0; i <= 10; i++ {
		q[i] = i
	}
	fmt.Println(q)
	r := make([]string, 3)
	names = [3]string{"Joe", "Joey", "Jamie"}
	for p, n := range names {
		r[p] = n
	}
	fmt.Println(r)
}

func main() {
	slicing()
}
