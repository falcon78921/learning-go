package main

import "fmt"

func goingDown() {
    for i := -1; i < 0; i-- {
        fmt.Println(i)
    }
}

func goingUp() {
    for i := 1; i > 0; i++ {
        fmt.Println(i)
    }
}

func main() {
    goingUp()
}
