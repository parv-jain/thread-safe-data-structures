package main

import (
    "fmt"
    "sync"
)

func main() {
    // Create a thread-safe map
    safeMap := new(sync.Map)

    var wg sync.WaitGroup

    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func(key int) {
            defer wg.Done()
            safeMap.Store("key", key)
        }(i)
    }

    wg.Wait()

    value, _ := safeMap.Load("key")
    fmt.Printf("Value: %d\n", value)
}
