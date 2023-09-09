package main

import (
	"fmt"
	"sync"
)

// SafeMap is a thread-safe map implementation.
type SafeMap struct {
	mu sync.Mutex
	data map[string]int
}

// NewSafeMap creates a new instance of SafeMap.
func NewSafeMap() *SafeMap {
	return &SafeMap{
		data: make(map[string]int),
	}
}

// Set sets a key-value pair in the map.
func (m *SafeMap) Set(key string, value int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[key] = value
}

// Get retrieves the value associated with a key.
func (m *SafeMap) Get(key string) (int, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	val, ok := m.data[key]
	return val, ok
}

// Delete removes a key-value pair from the map.
func (m *SafeMap) Delete(key string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.data, key)
}

// Size returns the number of key-value pairs in the map.
func (m *SafeMap) Size() int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return len(m.data)
}

func main() {
	// Create a new thread-safe map.
	safeMap := NewSafeMap()
	wg := sync.WaitGroup{}

	// Set key-value pairs concurrently.
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", i)
			safeMap.Set(key, i)
		}(i)
	}
	wg.Wait()

	// Get and print values concurrently.
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", i)
			val, ok := safeMap.Get(key)
			if ok {
				fmt.Printf("Key: %s, Value: %d\n", key, val)
			} else {
				fmt.Printf("Key: %s not found\n", key)
			}
		}(i)
	}
	wg.Wait()

	// Delete keys concurrently.
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", i)
			safeMap.Delete(key)
		}(i)
	}
	wg.Wait()

	// Print the final size of the map.
	size := safeMap.Size()
	fmt.Printf("Map size: %d\n", size)	
}
