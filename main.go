package main

import "fmt"

func main() {
    total := 0
    K := 9

    for i := 1; i <= K; i++ {
        if i%2 == 0 {
            total = total + (i * 2)
        } else {
            total = total - i
        }


        fmt.Println("Iteration:", i, "Total:", total)
    }

    fmt.Println("Final Output:", total)
}