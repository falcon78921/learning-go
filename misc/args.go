package main

import "fmt"
import "os"

func getArgs() []string {
	args := os.Args[1:]
	return args
}

func main() {
	fmt.Println(getArgs())
}
