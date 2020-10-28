package main

import "fmt"

func sameOld() {
    for i := 1; i <= 5; i++ {
        if i == 0 {
            fmt.Println("Zero")
        } else if i == 1 {
            fmt.Println("One")
        } else if i == 2 {
            fmt.Println("Two")
        } else if i == 3 {
            fmt.Println("Three")
        } else if i == 4 {
            fmt.Println("Four")
        }
    }
}

func switchTest() {
    for i := 1; i <= 5; i++ {
        switch i {
            case 0: fmt.Println("Zero")
            case 1: fmt.Println("One")
            case 2: fmt.Println("Two")
            case 3: fmt.Println("Three")
            case 4: fmt.Println("Four")
            case 5: fmt.Println("Five")
            default: fmt.Println("Unknown Number")
        }
    }
}

func main() {
    sameOld()
    switchTest()
}
