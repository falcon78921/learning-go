package main

import "fmt"

func main() {
    var scores [5]float64
    scores[0] = 98
    scores[1] = 93
    scores[2] = 77
    scores[3] = 82
    scores[4] = 83

    var total float64 = 0
    for i := 0; i < 5; i++ {
        total += scores[i]
    }
    fmt.Println(total / 5)
    improvedCalc()
}

func improvedCalc() {
    var scores [5]float64
    scores[0] = 98
    scores[1] = 93
    scores[2] = 77
    scores[3] = 82
    scores[4] = 83

    var total float64 = 0
    for i := 0; i < len(scores); i++ {
        total += scores[i]
    }
    fmt.Println(total / float64(len(scores)))
}
