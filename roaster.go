package main

import "fmt"

func main() {
    var N int
    var temperature float64

    ideal := 0
    overheated := 0
    totalTemperature := 0.0

    fmt.Print("Enter observation duration: ")
    fmt.Scan(&N)

    for i := 1; i <= N; i++ {
        fmt.Print("Enter temperature: ")
        fmt.Scan(&temperature)

        totalTemperature += temperature

        if temperature >= 195 && temperature <= 210 {
            ideal++
        } else if temperature > 210 {
            overheated++
        }
    }

    average := totalTemperature / float64(N)

    fmt.Println("Ideal Minutes:", ideal)
    fmt.Println("Overheated Minutes:", overheated)
    fmt.Println("Average Temperature:", average)
}