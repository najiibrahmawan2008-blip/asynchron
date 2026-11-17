package main

import "fmt"

func ComputePattern(n int) int {
    if n <= 1 {
        return 1
    }

    if n%2 == 1 {
        return n + ComputePattern(n-1)
    }

    return ComputePattern(n-2) * 2
}

func main() {
    result := ComputePattern(3)
    fmt.Println("Final Result:", result)
}