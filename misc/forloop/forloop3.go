package main

import "fmt"

func main() {
    for i := 1; i <= 10; i++ {
        fmt.Println("Hello World!")
        numbers()
    }
}

func numbers() {
    i := 1
    for i <= 10 {
        fmt.Println("Hello")
        i = i + 1
    }
}
