package main

import "fmt"

func fizzBuzzIf() {
    for i := 1; i <= 100; i++ {
        if i % 3 == 0 && i % 5 == 0 {
            fmt.Println(i)
            fmt.Println("FizzBuzz")
        } else if i % 3 == 0 {
            fmt.Println(i)
            fmt.Println("Fizz")
        } else if i % 5 == 0 {
            fmt.Println(i)
            fmt.Println("Buzz")
        }
    }
}

func main() {
    fizzBuzzIf()
}
