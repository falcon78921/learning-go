package main

import (
    "fmt"
    "net/http"
    "os"
    "io/ioutil"
)

var command string = os.Args[1]
var stock string = os.Args[2] 

func usage() {
    fmt.Println("stockfetcher: Fetch stock information from Yahoo Finance")
    fmt.Println("Usage:")
    fmt.Println("    --quote: Fetch stock quote for specified ticker")
}

func fetchStockInfo(stock string) {
    resp, err := http.Get("https://finance.yahoo.com/quote/"+ stock)

    if err != nil {
        fmt.Println("ERROR: Hmmm, is the connection stable?")
        os.Exit(1)
    } else {
        defer resp.Body.Close()
        body, err := ioutil.ReadAll(resp.Body)

        if err != nil {
            fmt.Println("ERROR: Cannot read response")
            os.Exit(1)
        } else {
            stockData := string(body)
            fmt.Printf("%s", stockData)
        }

    }

}

func main() {
    if command == "--quote" {
        fetchStockInfo(stock)
    } else if command == "--help" {
        usage()
    } else {
        usage()
        fmt.Println("ERROR: Please provide a valid option for stockfetcher.")
        os.Exit(1)
    }
}
