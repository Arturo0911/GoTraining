package main

import (
	"fmt"
	"sync"
	"time"
)

/*func main() {

	var data int
	go func() {
		data++
	}()
	time.Sleep(1 * time.Second)
	if data == 0 {
		fmt.Printf("The value is %v\n", data)
	}
	fmt.Printf("\nData printed %v\n\n", data)

}*/

/* Normal Example
func main() {
	var memoryAccess sync.Mutex
	var value int
	go func() {
		memoryAccess.Lock()
		value++
		memoryAccess.Unlock()
	}()

	memoryAccess.Lock()
	if value == 0 {
		fmt.Printf("\nData value %v\n\n", value)
	} else {
		fmt.Printf("\nData value %v\n\n", value)
	}
	memoryAccess.Unlock()
}*/

// Example with deadLocks
type value struct {
	mu    sync.Mutex
	value int
}

func main() {
	var wg sync.WaitGroup
	printSum := func(v1, v2 *value) {
		defer wg.Done()
		v1.mu.Lock()
		defer v1.mu.Unlock()

		time.Sleep(2 * time.Second)
		v2.mu.Lock()
		defer v2.mu.Unlock()
		fmt.Printf("\nsum of values is %v", v1.value+v2.value)
	}

	var a, b value
	wg.Add(2)
	go printSum(&a, &b)
	go printSum(&b, &a)
	wg.Wait()
}
