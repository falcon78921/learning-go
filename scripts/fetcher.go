package main

import "fmt"
import "net/http"
import "os"

func fetcher() {
    site := os.Args[1]
    response, err := http.Get(site)
    fmt.Println(response)
    if err != nil {
        fmt.Println("Houston, we have a problem...")
    }
}

func main() {
    fetcher()
}
