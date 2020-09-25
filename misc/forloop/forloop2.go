package main

import "os/exec"
import "fmt"
import "log"

func main() {
    for i := 1; i <= 25; i++ {
        out, err := exec.Command("date").Output()
        fmt.Printf("here is the date: %s", out)
        if err != nil {
            log.Fatal("no date?")
        }
    }
}
