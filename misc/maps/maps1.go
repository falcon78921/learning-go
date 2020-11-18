package main

import "fmt"

func main() {
    abigmap := make(map[string]string)
    abigmap["Founder"] = "Ray Kroc"
    abigmap["Portrayed By"] = "Michael Keaton"
    abigmap["Film Name"] = "The Founder"
    abigmap["Release Year"] = "2016"
    fmt.Println(abigmap)
    fmt.Println(abigmap["Founder"])
    delete(abigmap, "Founder")
    fmt.Println(abigmap)
}
