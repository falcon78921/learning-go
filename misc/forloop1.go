package main

import "fmt"

func otherWay() {
	for i := 1; i <= -15; i-- {
	    fmt.Println(i)
	}
}

func main() {
	for i := 1; i <= 15; i++ {
	    fmt.Println(i)
	}
	otherWay()
}
