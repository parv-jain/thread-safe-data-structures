package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    var resource1 sync.Mutex
    var resource2 sync.Mutex

    // Function to run Goroutine 1 and Goroutine 2
    run := func() {
        // Goroutine 1
        go func() {
            fmt.Println("Goroutine 1: Trying to lock resource1...")
            resource1.Lock()
            defer resource1.Unlock()
            fmt.Println("Goroutine 1: Locked resource1")

            // Introduce a delay to make the deadlock more likely
            time.Sleep(time.Millisecond * 100)

            fmt.Println("Goroutine 1: Trying to lock resource2...")
            resource2.Lock()
            defer resource2.Unlock()
            fmt.Println("Goroutine 1: Locked resource2")
        }()

        // Goroutine 2
        go func() {
            fmt.Println("Goroutine 2: Trying to lock resource2...")
            resource2.Lock()
            defer resource2.Unlock()
            fmt.Println("Goroutine 2: Locked resource2")

            // Introduce a delay to make the deadlock more likely
            time.Sleep(time.Millisecond * 100)

            fmt.Println("Goroutine 2: Trying to lock resource1...")
            resource1.Lock()
            defer resource1.Unlock()
            fmt.Println("Goroutine 2: Locked resource1")
        }()
    }

    // Run Goroutine 1 and Goroutine 2 in a loop
    for {
        run()
        // Introduce a delay between iterations
        time.Sleep(time.Second)
    }
}
