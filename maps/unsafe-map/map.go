package main

import (
	"fmt"
)

func main() {
	unsafeMap := make(map[string]int)
	
	for i := 0; i < 100; i++ {
		go func(key int) {
			unsafeMap["key"] = key
		}(i)
	}
	
	value := unsafeMap["key"]
	fmt.Printf("Value: %d\n", value)
}

