package main

import (
    "fmt"
    "strconv"
)

func main() {
    var result = 0
    for i := 1; i < 1000; i++ {
        if i%3 == 0 || i%5 == 0 {
            result += i
        }
    }

    fmt.Println("result: " + strconv.Itoa(result))
}