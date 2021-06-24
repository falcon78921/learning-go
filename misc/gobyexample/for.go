package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
		j := i + 3
		fmt.Println(j)
		if j == 12 {
			fmt.Println("REACHED!")
			break
		}
	}

	for l := 0; l <= 10; l++ {
		fmt.Println(l)
		m := l + 2
		fmt.Println(m)
		if m == 11 {
			fmt.Println("REACHED!")
			continue
		}
	}
}
