package main

import "fmt"

func readAge(name, age string) string {
    return name + " is " + age + " years old!"
}

func main() {
    output := readAge("Joey", "24")
    fmt.Println(output)
}
