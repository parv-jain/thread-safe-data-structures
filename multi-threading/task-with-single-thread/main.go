package main

import (
    "fmt"
    "time"
)

func calculateSumSequential(N int) int {
    start := time.Now()
    sum := 0
    for i := 1; i <= N; i++ {
        sum += i
    }
    elapsed := time.Since(start)
    fmt.Printf("Sum (Sequential): %d, Time: %s\n", sum, elapsed)
    return sum
}

func main() {
    N := 1000000000 // Sum all numbers from 1 to 1 billion
    calculateSumSequential(N)
}
