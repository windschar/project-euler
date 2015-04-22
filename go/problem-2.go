package main

import (
    "fmt"
    "strconv"
)

func main() {
    var maxValue = 4000000
    var sum = 0

    for a, b := 1, 2; b <= maxValue; a, b = b, a+b {
        if b%2 == 0 {
            sum += b
        }
    }

    fmt.Println("sum: " + strconv.Itoa(sum))
}