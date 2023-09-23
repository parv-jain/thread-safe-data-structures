package main

import (
    "fmt"
    "sync"
    "time"
)

func calculateSumParallel(N, numWorkers int) int {
    start := time.Now()
    chunkSize := N / numWorkers
    sum := 0
    var wg sync.WaitGroup

    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go func(start, end int) {
            defer wg.Done()
            localSum := 0
            for j := start; j <= end; j++ {
                localSum += j
            }
            sum += localSum
        }(i*chunkSize+1, (i+1)*chunkSize)
    }

    wg.Wait()
    elapsed := time.Since(start)
    fmt.Printf("Sum (Parallel): %d, Time: %s\n", sum, elapsed)
    return sum
}

func main() {
    N := 1000000000 // Sum all numbers from 1 to 1 billion
    numWorkers := 4 // Number of worker goroutines
    calculateSumParallel(N, numWorkers)
}
