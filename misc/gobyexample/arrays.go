package main

import "fmt"

var storage [5]int
var multiDimensionalStorage [2][4]int

func createArray() [5]int {
	for i := 0; i <= 4; i++ {
		storage[i] = i
	}

	return storage
}

func multiDimensional() [2][4]int {
	for i := 0; i <= 1; i++ {
		for j := 0; j <= 3; j++ {
			multiDimensionalStorage[i][j] = i + j
		}
	}

	return multiDimensionalStorage
}

func main() {
	moreStorage := [5]int{6, 7, 8, 9, 10}
	p := createArray()
	q := multiDimensional()
	fmt.Println(p)
	fmt.Println(len(p))
	fmt.Println(moreStorage)
	fmt.Println(len(moreStorage))
	fmt.Println(q)
}
